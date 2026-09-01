package paginate_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/stackitcloud/stackit-sdk-go/experimental/paginate"
)

type testItem struct {
	ID int `json:"id"`
}

type testResponse struct {
	Items         []testItem `json:"items"`
	NextPageToken string     `json:"nextPageToken,omitempty"`
}

func (r *testResponse) GetItems() []testItem     { return r.Items }
func (r *testResponse) GetNextPageToken() string { return r.NextPageToken }

type testRequest struct {
	ctx       context.Context
	client    *http.Client
	endpoint  string
	pageSize  int32
	pageToken string
}

func (r testRequest) Context(ctx context.Context) testRequest {
	r.ctx = ctx
	return r
}

func (r testRequest) PageSize(pageSize int32) testRequest {
	r.pageSize = pageSize
	return r
}

func (r testRequest) PageToken(pageToken string) testRequest {
	r.pageToken = pageToken
	return r
}

func (r testRequest) Execute() (*testResponse, error) {
	u, err := url.Parse(r.endpoint)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	if r.pageSize > 0 {
		query.Set("pageSize", strconv.FormatInt(int64(r.pageSize), 10))
	}
	if r.pageToken != "" {
		query.Set("pageToken", r.pageToken)
	}
	u.RawQuery = query.Encode()

	ctx := r.ctx
	if ctx == nil {
		ctx = context.Background()
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected HTTP status %s", resp.Status)
	}

	var result testResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func TestAllPaginatesHTTPListOperation(t *testing.T) {
	var queries []url.Values
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		queries = append(queries, r.URL.Query())
		pages := map[string]testResponse{
			"":       {Items: []testItem{{1}, {2}, {3}}, NextPageToken: "page-a"},
			"page-a": {Items: []testItem{{4}, {5}, {6}}, NextPageToken: "page-b"},
			"page-b": {Items: []testItem{{7}}},
		}
		page, ok := pages[r.URL.Query().Get("pageToken")]
		if !ok {
			http.Error(w, "unknown page token", http.StatusBadRequest)
			return
		}
		_ = json.NewEncoder(w).Encode(page)
	}))
	defer server.Close()

	got, err := paginate.All(newTestRequest(server), paginate.WithPageSize(3))
	if err != nil {
		t.Fatalf("All() error = %v", err)
	}
	want := []testItem{{1}, {2}, {3}, {4}, {5}, {6}, {7}}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("All() mismatch (-want +got):\n%s", diff)
	}

	wantQueries := []url.Values{
		{"pageSize": {"3"}},
		{"pageSize": {"3"}, "pageToken": {"page-a"}},
		{"pageSize": {"3"}, "pageToken": {"page-b"}},
	}
	if diff := cmp.Diff(wantQueries, queries); diff != "" {
		t.Errorf("queries mismatch (-want +got):\n%s", diff)
	}
}

func TestAllLimitAdjustsLastPageSize(t *testing.T) {
	var queries []url.Values
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		queries = append(queries, r.URL.Query())
		if r.URL.Query().Get("pageToken") == "" {
			_ = json.NewEncoder(w).Encode(testResponse{
				Items: []testItem{{1}, {2}, {3}, {4}}, NextPageToken: "next",
			})
			return
		}
		_ = json.NewEncoder(w).Encode(testResponse{
			Items: []testItem{{5}, {6}}, NextPageToken: "not-requested",
		})
	}))
	defer server.Close()

	got, err := paginate.All(newTestRequest(server), paginate.WithPageSize(4), paginate.WithLimit(6))
	if err != nil {
		t.Fatalf("All() error = %v", err)
	}
	want := []testItem{{1}, {2}, {3}, {4}, {5}, {6}}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("All() mismatch (-want +got):\n%s", diff)
	}
	wantQueries := []url.Values{
		{"pageSize": {"4"}},
		{"pageSize": {"2"}, "pageToken": {"next"}},
	}
	if diff := cmp.Diff(wantQueries, queries); diff != "" {
		t.Errorf("queries mismatch (-want +got):\n%s", diff)
	}
}

func TestItemsIsLazyAndStopsWithConsumer(t *testing.T) {
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requests++
		_ = json.NewEncoder(w).Encode(testResponse{
			Items: []testItem{{requests*2 - 1}, {requests * 2}}, NextPageToken: fmt.Sprintf("page-%d", requests+1),
		})
	}))
	defer server.Close()

	sequence := paginate.Items(newTestRequest(server), paginate.WithPageSize(2))
	if requests != 0 {
		t.Fatalf("Items made %d requests before iteration", requests)
	}

	var got []testItem
	for item, err := range sequence {
		if err != nil {
			t.Fatalf("Items() error = %v", err)
		}
		got = append(got, item)
		if len(got) == 2 {
			break
		}
	}
	want := []testItem{{1}, {2}}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Items() mismatch (-want +got):\n%s", diff)
	}
	if requests != 1 {
		t.Errorf("request count = %d, want 1", requests)
	}
}

func TestAllContinuesAfterEmptyNonFinalPage(t *testing.T) {
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requests++
		if requests == 1 {
			_ = json.NewEncoder(w).Encode(testResponse{NextPageToken: "next"})
			return
		}
		_ = json.NewEncoder(w).Encode(testResponse{Items: []testItem{{1}}})
	}))
	defer server.Close()

	got, err := paginate.All(newTestRequest(server))
	if err != nil {
		t.Fatalf("All() error = %v", err)
	}
	want := []testItem{{1}}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("All() mismatch (-want +got):\n%s", diff)
	}
	if requests != 2 {
		t.Errorf("request count = %d, want 2", requests)
	}
}

func TestAllReturnsPartialItemsAndHTTPError(t *testing.T) {
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requests++
		if requests == 1 {
			_ = json.NewEncoder(w).Encode(testResponse{Items: []testItem{{1}}, NextPageToken: "next"})
			return
		}
		http.Error(w, "failed", http.StatusInternalServerError)
	}))
	defer server.Close()

	got, err := paginate.All(newTestRequest(server))
	if err == nil {
		t.Fatal("All() error = nil, want an error")
	}
	want := []testItem{{1}}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("All() partial result mismatch (-want +got):\n%s", diff)
	}
	if want := "paginate: fetch page 2"; len(err.Error()) < len(want) || err.Error()[:len(want)] != want {
		t.Errorf("All() error = %q, want prefix %q", err, want)
	}
}

func TestAllMaxPages(t *testing.T) {
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requests++
		_ = json.NewEncoder(w).Encode(testResponse{
			Items: []testItem{{requests}}, NextPageToken: fmt.Sprintf("next-%d", requests),
		})
	}))
	defer server.Close()

	got, err := paginate.All(newTestRequest(server), paginate.WithMaxPages(2))
	if err != nil {
		t.Fatalf("All() error = %v", err)
	}
	want := []testItem{{1}, {2}}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("All() mismatch (-want +got):\n%s", diff)
	}
	if requests != 2 {
		t.Errorf("request count = %d, want 2", requests)
	}
}

func TestAllRejectsInvalidOptionsWithoutRequest(t *testing.T) {
	tests := []struct {
		name string
		opt  paginate.Option
	}{
		{name: "negative page size", opt: paginate.WithPageSize(-1)},
		{name: "negative limit", opt: paginate.WithLimit(-1)},
		{name: "negative maximum pages", opt: paginate.WithMaxPages(-1)},
		{name: "nil option", opt: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requests := 0
			server := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {
				requests++
			}))
			defer server.Close()

			_, err := paginate.All(newTestRequest(server), tt.opt)
			if err == nil {
				t.Fatal("All() error = nil, want an error")
			}
			if requests != 0 {
				t.Errorf("request count = %d, want 0", requests)
			}
		})
	}
}

func TestAllZeroLimitsDoNotRequest(t *testing.T) {
	for _, opt := range []paginate.Option{paginate.WithLimit(0), paginate.WithMaxPages(0)} {
		requests := 0
		server := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {
			requests++
		}))

		got, err := paginate.All(newTestRequest(server), opt)
		server.Close()
		if err != nil {
			t.Fatalf("All() error = %v", err)
		}
		if len(got) != 0 {
			t.Errorf("All() = %v, want no items", got)
		}
		if requests != 0 {
			t.Errorf("request count = %d, want 0", requests)
		}
	}
}

func TestAllOptionPrecedence(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_ = json.NewEncoder(w).Encode(testResponse{
			Items: []testItem{{1}, {2}, {3}, {4}, {5}},
		})
	}))
	defer server.Close()

	// Later option overrides earlier option
	got, err := paginate.All(newTestRequest(server), paginate.WithLimit(1), paginate.WithLimit(3))
	if err != nil {
		t.Fatalf("All() error = %v", err)
	}
	want := []testItem{{1}, {2}, {3}}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("All() mismatch (-want +got):\n%s", diff)
	}
}

func TestItemsStopsOnRepeatedPageToken(t *testing.T) {
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requests++
		_ = json.NewEncoder(w).Encode(testResponse{NextPageToken: "same-token"})
	}))
	defer server.Close()

	_, err := paginate.All(newTestRequest(server))
	if err == nil {
		t.Fatal("All() error = nil, want an error")
	}
	if requests != 2 {
		t.Errorf("request count = %d, want 2", requests)
	}
}

func TestAllServerCapsPageSize(t *testing.T) {
	var queries []url.Values
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		queries = append(queries, r.URL.Query())
		pages := map[string]testResponse{
			"":       {Items: []testItem{{1}, {2}}, NextPageToken: "page-2"},
			"page-2": {Items: []testItem{{3}, {4}}, NextPageToken: "page-3"},
			"page-3": {Items: []testItem{{5}}},
		}
		page, ok := pages[r.URL.Query().Get("pageToken")]
		if !ok {
			http.Error(w, "unknown page token", http.StatusBadRequest)
			return
		}
		_ = json.NewEncoder(w).Encode(page)
	}))
	defer server.Close()

	got, err := paginate.All(newTestRequest(server), paginate.WithPageSize(3))
	if err != nil {
		t.Fatalf("All() error = %v", err)
	}
	want := []testItem{{1}, {2}, {3}, {4}, {5}}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("All() mismatch (-want +got):\n%s", diff)
	}

	wantQueries := []url.Values{
		{"pageSize": {"3"}},
		{"pageSize": {"3"}, "pageToken": {"page-2"}},
		{"pageSize": {"3"}, "pageToken": {"page-3"}},
	}
	if diff := cmp.Diff(wantQueries, queries); diff != "" {
		t.Errorf("queries mismatch (-want +got):\n%s", diff)
	}
}

func TestItemsContextCanceledMidIteration(t *testing.T) {
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		pages := map[string]testResponse{
			"":       {Items: []testItem{{1}, {2}}, NextPageToken: "page-2"},
			"page-2": {Items: []testItem{{3}, {4}}},
		}
		page, ok := pages[r.URL.Query().Get("pageToken")]
		if !ok {
			http.Error(w, "unknown page token", http.StatusBadRequest)
			return
		}
		_ = json.NewEncoder(w).Encode(page)
	}))
	defer server.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var got []testItem
	var iterErr error
	for item, err := range paginate.Items(newTestRequest(server).Context(ctx)) {
		if err != nil {
			iterErr = err
			break
		}
		got = append(got, item)
		if len(got) == 2 {
			cancel()
		}
	}

	if iterErr == nil {
		t.Fatal("Items() error = nil, want context canceled error")
	}
	if !errors.Is(iterErr, context.Canceled) {
		t.Errorf("Items() error = %v, want context.Canceled", iterErr)
	}
	want := []testItem{{1}, {2}}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Items() mismatch (-want +got):\n%s", diff)
	}
	if requests != 1 {
		t.Errorf("request count = %d, want 1", requests)
	}
}

func newTestRequest(server *httptest.Server) testRequest {
	return testRequest{ctx: context.Background(), client: server.Client(), endpoint: server.URL}
}
