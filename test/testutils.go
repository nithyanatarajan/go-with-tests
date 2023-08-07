package test

import (
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, got, want string) {
	t.Helper() // When failed the line number reported will be in called function

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertEqualInt(t *testing.T, got, want int) {
	t.Helper() // When failed the line number reported will be in called function

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertEqualIntArray(t *testing.T, got, want []int) {
	t.Helper() // When failed the line number reported will be in called function

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
