package utils

import (
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
