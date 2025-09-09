package utils

import (
	"encoding/base64"
	"testing"
)

func TestPtrConversions(t *testing.T) {
	s := "test"
	sPrt := Ptr(s)
	if *sPrt != s {
		t.Fatalf("conversion failed")
	}
}

func TestContainsString(t *testing.T) {
	sl := []string{"a", "b"}
	if Contains(sl, "c") {
		t.Fatalf("Should not be contained")
	}
	if !Contains(sl, "a") {
		t.Fatalf("Should be contained")
	}
	if Contains([]string{}, "a") {
		t.Fatalf("Should not be contained")
	}
}

func TestContainsInt(t *testing.T) {
	sl := []int{11, 22}
	if Contains(sl, 33) {
		t.Fatalf("Should not be contained")
	}
	if !Contains(sl, 11) {
		t.Fatalf("Should be contained")
	}
	if Contains([]int{}, 11) {
		t.Fatalf("Should not be contained")
	}
}

// Test struct for YAML conversion testing
type TestServer struct {
	Name     string  `json:"name"`
	UserData *[]byte `json:"userData,omitempty"`
	Data     []byte  `json:"data,omitempty"`
}

func TestConvertByteArraysToBase64(t *testing.T) {
	tests := []struct {
		name           string
		input          interface{}
		expectedFields map[string]interface{}
		expectError    bool
	}{
		{
			name:        "nil input",
			input:       nil,
			expectError: false,
		},
		{
			name: "normal case with byte arrays",
			input: TestServer{
				Name:     "test-server",
				UserData: func() *[]byte { b := []byte("hello world"); return &b }(),
				Data:     []byte("test data"),
			},
			expectedFields: map[string]interface{}{
				"name":     "test-server",
				"userData": base64.StdEncoding.EncodeToString([]byte("hello world")),
				"data":     base64.StdEncoding.EncodeToString([]byte("test data")),
			},
			expectError: false,
		},
		{
			name: "nil pointer case",
			input: TestServer{
				Name:     "test-server",
				UserData: nil,
				Data:     []byte("test"),
			},
			expectedFields: map[string]interface{}{
				"name": "test-server",
				"data": base64.StdEncoding.EncodeToString([]byte("test")),
			},
			expectError: false,
		},
		{
			name: "empty byte array case",
			input: TestServer{
				Name: "test-server",
				Data: []byte{},
			},
			expectedFields: map[string]interface{}{
				"name": "test-server",
				// Note: empty byte arrays are omitted from JSON output
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertByteArraysToBase64(tt.input)

			if tt.expectError {
				if err == nil {
					t.Fatalf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("Expected no error, got: %v", err)
			}

			// Special case for nil input
			if tt.input == nil {
				if result != nil {
					t.Fatalf("Expected nil result for nil input, got: %v", result)
				}
				return
			}

			// Check expected fields
			for fieldName, expectedValue := range tt.expectedFields {
				if actualValue, exists := result[fieldName]; !exists {
					// For empty byte arrays, the field might be omitted from JSON
					if fieldName == "data" && expectedValue == nil {
						t.Logf("Empty byte array was omitted from JSON output, which is expected")
						continue
					}
					t.Fatalf("Expected field %s to exist, but it doesn't", fieldName)
				} else {
					if actualValue != expectedValue {
						t.Fatalf("Expected field %s to be %v, got %v", fieldName, expectedValue, actualValue)
					}
				}
			}

			// Check that no unexpected fields exist (except for omitted empty byte arrays)
			for fieldName, actualValue := range result {
				if _, expected := tt.expectedFields[fieldName]; !expected {
					// Allow data field to exist even if not in expectedFields (for empty byte array case)
					if fieldName == "data" && tt.name == "empty byte array case" {
						continue
					}
					t.Fatalf("Unexpected field %s with value %v", fieldName, actualValue)
				}
			}
		})
	}
}

func TestConvertByteArraysToBase64Recursive(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string // check if []byte was converted to base64 string
	}{
		{"nil", nil, ""},
		{"map with []byte", map[string]interface{}{"data": []byte("hello")}, "aGVsbG8="},
		{"slice with []byte", []interface{}{[]byte("test")}, "dGVzdA=="},
		{"nested map", map[string]interface{}{"level": map[string]interface{}{"data": []byte("nested")}}, "bmVzdGVk"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			convertByteArraysToBase64Recursive(tt.input)

			if tt.expected == "" {
				return
			}

			// check if base64 string in the result
			found := findBase64String(tt.input)
			if found != tt.expected {
				t.Fatalf("Expected %s, got %s", tt.expected, found)
			}
		})
	}
}

// helper to find base64 string in interface{}
func findBase64String(data interface{}) string {
	switch v := data.(type) {
	case map[string]interface{}:
		for _, val := range v {
			if str := findBase64String(val); str != "" {
				return str
			}
		}
	case []interface{}:
		for _, val := range v {
			if str := findBase64String(val); str != "" {
				return str
			}
		}
	case string:
		if len(v) > 0 && v != "hello" && v != "test" && v != "nested" {
			return v
		}
	}
	return ""
}
