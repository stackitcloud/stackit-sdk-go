// Package paginate provides lazy and eager pagination for list operations that
// follow AIP-158.
package paginate

import (
	"fmt"
	"iter"
	"math"
)

// Response defines the methods needed for pagination. The result of Request.Execute() must implement this interface.
type Response[Item any] interface {
	GetItems() []Item
	GetNextPageToken() string
}

// Request defines the methods needed for pagination on te request side.
type Request[Self, Resp any] interface {
	PageSize(pageSize int32) Self
	PageToken(pageToken string) Self
	Execute() (Resp, error)
}

type Option func(*options) error

type options struct {
	pageSize int32
	limit    *int
	maxPages *int
}

// WithPageSize sets the preferred number of items requested per page. The
// service may return fewer items. A value of zero does not override the page
// size already set on the supplied request.
func WithPageSize(pageSize int32) Option {
	return func(o *options) error {
		if pageSize < 0 {
			return fmt.Errorf("page size must not be negative: %d", pageSize)
		}
		o.pageSize = pageSize
		return nil
	}
}

// WithLimit sets the maximum number of items yielded across all pages. A zero
// limit performs no requests and yields no items.
func WithLimit(limit int) Option {
	return func(o *options) error {
		if limit < 0 {
			return fmt.Errorf("limit must not be negative: %d", limit)
		}
		o.limit = &limit
		return nil
	}
}

// WithMaxPages sets the maximum number of pages fetched. A zero maximum
// performs no requests. Reaching the maximum is not an error.
func WithMaxPages(maxPages int) Option {
	return func(o *options) error {
		if maxPages < 0 {
			return fmt.Errorf("maximum pages must not be negative: %d", maxPages)
		}
		o.maxPages = &maxPages
		return nil
	}
}

// Items returns a lazy iterator over the items of pageable ist operation.
// No request is made until the iterator is consumed. Each successful item is
// yielded with a nil error. If a request or option fails, the error is yielded
// once with the zero value of Item and iteration stops.
//
// Iteration also stops when the consumer returns false, the configured item or
// page limit is reached, or the service returns an empty next page token.
func Items[
	Item any,
	Resp Response[Item],
	Req Request[Req, Resp],
](request Req, opts ...Option) iter.Seq2[Item, error] {
	return func(yield func(Item, error) bool) {
		var zero Item

		cfg, err := applyOptions(opts)
		if err != nil {
			yield(zero, fmt.Errorf("paginate: invalid option: %w", err))
			return
		}
		if cfg.limit != nil && *cfg.limit == 0 {
			return
		}
		if cfg.maxPages != nil && *cfg.maxPages == 0 {
			return
		}

		seenTokens := make(map[string]struct{})
		itemCount := 0
		pageCount := 0

		for {
			pageSize := cfg.pageSize
			if cfg.limit != nil {
				remaining := *cfg.limit - itemCount
				if remaining <= 0 {
					return
				}
				if pageSize == 0 || int64(remaining) < int64(pageSize) {
					if remaining > math.MaxInt32 {
						pageSize = math.MaxInt32
					} else {
						pageSize = int32(remaining)
					}
				}
			}
			if pageSize > 0 {
				request = request.PageSize(pageSize)
			}

			response, err := request.Execute()
			if err != nil {
				yield(zero, fmt.Errorf("paginate: fetch page %d: %w", pageCount+1, err))
				return
			}
			pageCount++

			items := response.GetItems()
			if cfg.limit != nil {
				remaining := *cfg.limit - itemCount
				if len(items) > remaining {
					items = items[:remaining]
				}
			}
			for _, item := range items {
				itemCount++
				if !yield(item, nil) {
					return
				}
			}

			if cfg.limit != nil && itemCount >= *cfg.limit {
				return
			}

			nextPageToken := response.GetNextPageToken()
			if nextPageToken == "" {
				return
			}
			if cfg.maxPages != nil && pageCount >= *cfg.maxPages {
				return
			}
			if _, exists := seenTokens[nextPageToken]; exists {
				yield(zero, fmt.Errorf("paginate: page %d returned an already used next page token", pageCount))
				return
			}
			seenTokens[nextPageToken] = struct{}{}
			request = request.PageToken(nextPageToken)
		}
	}
}

// All retrieves and returns all items yielded by Items. If pagination fails,
// All returns the items retrieved before the failure together with the error.
func All[
	Item any,
	Resp Response[Item],
	Req Request[Req, Resp],
](request Req, opts ...Option) ([]Item, error) {
	var items []Item
	for item, err := range Items(request, opts...) {
		if err != nil {
			return items, err
		}
		items = append(items, item)
	}
	return items, nil
}

func applyOptions(opts []Option) (options, error) {
	var cfg options
	for i, opt := range opts {
		if opt == nil {
			return options{}, fmt.Errorf("option %d is nil", i+1)
		}
		if err := opt(&cfg); err != nil {
			return options{}, err
		}
	}
	return cfg, nil
}
