package utils

import (
	"reflect"
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

func TestEnumSliceToStringSlice(t *testing.T) {
	type TestEnum string

	const TESTENUM_CREATING TestEnum = "CREATING"
	const TESTENUM_ACTIVE TestEnum = "ACTIVE"
	const TESTENUM_UPDATING TestEnum = "UPDATING"
	const TESTENUM_DELETING TestEnum = "DELETING"
	const TESTENUM_ERROR TestEnum = "ERROR"

	type args[T interface{ ~string }] struct {
		inputSlice []T
	}
	type test[T interface{ ~string }] struct {
		name string
		args args[T]
		want []string
	}
	tests := []test[TestEnum]{
		{
			name: "default",
			args: args[TestEnum]{
				inputSlice: []TestEnum{
					TESTENUM_CREATING,
					TESTENUM_ACTIVE,
					TESTENUM_UPDATING,
					TESTENUM_DELETING,
					TESTENUM_ERROR,
				},
			},
			want: []string{"CREATING", "ACTIVE", "UPDATING", "DELETING", "ERROR"},
		},
		{
			name: "empty input slice",
			args: args[TestEnum]{
				inputSlice: []TestEnum{},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EnumSliceToStringSlice(tt.args.inputSlice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EnumSliceToStringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
