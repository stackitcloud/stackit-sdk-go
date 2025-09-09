package utils

import (
	"encoding/base64"
	"encoding/json"
)

// Ptr Returns the pointer to any type T
func Ptr[T any](v T) *T {
	return &v
}

func Contains[T comparable](slice []T, element T) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}

// ConvertByteArraysToBase64 converts a struct to a map with byte arrays converted to base64 strings.
// This is useful for serialization formats like YAML where []byte fields should be displayed
// as base64 strings instead of byte arrays. Works with any struct containing []byte fields.
func ConvertByteArraysToBase64(obj interface{}) (map[string]interface{}, error) {
	if obj == nil {
		return nil, nil
	}

	// Marshal to JSON first to get the correct structure
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return nil, err
	}

	// Convert byte arrays to base64 strings
	convertByteArraysToBase64(result)

	return result, nil
}

// convertByteArraysToBase64 recursively converts []byte values to base64 strings
func convertByteArraysToBase64(data interface{}) {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			if byteArray, ok := value.([]byte); ok {
				v[key] = base64.StdEncoding.EncodeToString(byteArray)
			} else {
				convertByteArraysToBase64(value)
			}
		}
	case []interface{}:
		for i, value := range v {
			if byteArray, ok := value.([]byte); ok {
				v[i] = base64.StdEncoding.EncodeToString(byteArray)
			} else {
				convertByteArraysToBase64(value)
			}
		}
	}
}
