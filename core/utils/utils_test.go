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
	// Test with nil input
	result, err := ConvertByteArraysToBase64(nil)
	if err != nil {
		t.Fatalf("Expected no error for nil input, got: %v", err)
	}
	if result != nil {
		t.Fatalf("Expected nil result for nil input, got: %v", result)
	}

	// Test with server containing byte arrays
	userData := []byte("hello world")
	data := []byte("test data")

	server := TestServer{
		Name:     "test-server",
		UserData: &userData,
		Data:     data,
	}

	result1, err := ConvertByteArraysToBase64(server)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	// Check that UserData is converted to base64
	if userDataStr, ok := result1["userData"].(string); !ok {
		t.Fatalf("Expected userData to be a string, got: %T", result1["userData"])
	} else {
		expected := base64.StdEncoding.EncodeToString(userData)
		if userDataStr != expected {
			t.Fatalf("Expected userData to be %s, got: %s", expected, userDataStr)
		}
	}

	// Check that Data is converted to base64
	if dataStr, ok := result1["data"].(string); !ok {
		t.Fatalf("Expected data to be a string, got: %T", result1["data"])
	} else {
		expected := base64.StdEncoding.EncodeToString(data)
		if dataStr != expected {
			t.Fatalf("Expected data to be %s, got: %s", expected, dataStr)
		}
	}

	// Check that other fields remain unchanged
	if result1["name"] != "test-server" {
		t.Fatalf("Expected name to be 'test-server', got: %v", result1["name"])
	}
}

func TestConvertByteArraysToBase64WithNilPointer(t *testing.T) {
	// Test with nil pointer to byte array
	server := TestServer{
		Name:     "test-server",
		UserData: nil, // nil pointer
		Data:     []byte("test"),
	}

	result2, err := ConvertByteArraysToBase64(server)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	// Check that nil pointer is handled correctly
	if result2["userData"] != nil {
		t.Fatalf("Expected userData to be nil, got: %v", result2["userData"])
	}

	// Check that non-nil byte array is still converted
	if dataStr, ok := result2["data"].(string); !ok {
		t.Fatalf("Expected data to be a string, got: %T", result2["data"])
	} else {
		expected := base64.StdEncoding.EncodeToString([]byte("test"))
		if dataStr != expected {
			t.Fatalf("Expected data to be %s, got: %s", expected, dataStr)
		}
	}
}

func TestConvertByteArraysToBase64WithEmptyByteArray(t *testing.T) {
	// Test with empty byte array
	emptyData := []byte{}

	server := TestServer{
		Name: "test-server",
		Data: emptyData,
	}

	result3, err := ConvertByteArraysToBase64(server)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	// Check that empty byte array is converted to empty base64 string
	// Note: empty byte arrays might be omitted from JSON, so we check if it exists
	if dataStr, exists := result3["data"]; !exists {
		// Empty byte array was omitted from JSON, which is expected behavior
		t.Logf("Empty byte array was omitted from JSON output, which is expected")
	} else if dataStr, ok := dataStr.(string); !ok {
		t.Fatalf("Expected data to be a string, got: %T", result3["data"])
	} else {
		expected := base64.StdEncoding.EncodeToString(emptyData)
		if dataStr != expected {
			t.Fatalf("Expected data to be %s, got: %s", expected, dataStr)
		}
		if dataStr != "" {
			t.Fatalf("Expected empty byte array to be converted to empty string, got: %s", dataStr)
		}
	}
}
