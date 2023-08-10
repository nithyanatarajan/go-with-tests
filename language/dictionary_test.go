package language_test

import (
	"testing"

	"github.com/nithyanatarajan/go-with-tests/language"
	. "github.com/nithyanatarajan/go-with-tests/test"
)

func TestSearch(t *testing.T) {
	dictionary := language.Dictionary{"test": "this is a test"}

	t.Run("should return value given key", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "this is a test"

		AssertNoError(t, err)
		AssertEqual(t, got, want)
	})

	t.Run("should return error when given key not present", func(t *testing.T) {
		got, err := dictionary.Search("unknown")
		want := ""

		AssertError(t, err)
		AssertEqual(t, err, language.ErrNotFound)
		AssertEqual(t, got, want)
	})
}

func TestAdd(t *testing.T) {
	dictionary := language.Dictionary{"test": "this is a test"}
	t.Run("add a word", func(t *testing.T) {
		err := dictionary.Add("apple", "it is a fruit")
		AssertNoError(t, err)

		got, err := dictionary.Search("apple")
		want := "it is a fruit"

		AssertNoError(t, err)
		AssertEqual(t, got, want)
	})

	t.Run("should return error when given key is already present",
		func(t *testing.T) {
			err := dictionary.Add("test", "it is a fruit")
			AssertError(t, err)
			AssertEqual(t, err, language.ErrWordAlreadyPresent)

			got, err := dictionary.Search("test")
			want := "this is a test"

			AssertNoError(t, err)
			AssertEqual(t, got, want)
		})
}

func TestUpdate(t *testing.T) {
	dictionary := language.Dictionary{"test": "this is a test"}
	t.Run("should update word", func(t *testing.T) {
		err := dictionary.Update("test", "this is a updated test")
		AssertNoError(t, err)

		got, err := dictionary.Search("test")
		want := "this is a updated test"

		AssertNoError(t, err)
		AssertEqual(t, got, want)
	})

	t.Run("should return error when given key is not present",
		func(t *testing.T) {
			err := dictionary.Update("apple", "it is a fruit")

			AssertError(t, err)
			AssertEqual(t, err, language.ErrWordNotPresent)

			got, err := dictionary.Search("apple")
			want := ""

			AssertError(t, err)
			AssertEqual(t, err, language.ErrNotFound)
			AssertEqual(t, got, want)
		})
}

func TestDelete(t *testing.T) {
	dictionary := language.Dictionary{"test": "this is a test"}
	t.Run("should delete word", func(t *testing.T) {
		err := dictionary.Delete("test")
		AssertNoError(t, err)

		got, err := dictionary.Search("test")
		want := ""

		AssertError(t, err)
		AssertEqual(t, err, language.ErrNotFound)
		AssertEqual(t, got, want)
	})

	t.Run("should return error when given key is not present",
		func(t *testing.T) {
			err := dictionary.Delete("apple")

			AssertError(t, err)
			AssertEqual(t, err, language.ErrNotFound)
		})
}
