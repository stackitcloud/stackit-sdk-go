package utils

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
