package test

import (
	"reflect"
	"testing"
)

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
