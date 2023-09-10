package test

import (
	"reflect"
	"testing"
)

func doesContain(haystack []string, needle string) bool {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	return contains
}
func AssertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()
	if !doesContain(haystack, needle) {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}

func AssertEqual(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("wanted an error but didn't get any")
	}
}

func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("error not expected but got one")
	}
}
