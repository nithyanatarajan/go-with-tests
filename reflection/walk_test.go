package reflection_test

import (
	"testing"

	"github.com/nithyanatarajan/go-with-tests/reflection"
	. "github.com/nithyanatarajan/go-with-tests/test"
)

type profile struct {
	age      int
	location string
}

var cases = []struct {
	name  string
	input any
	want  []string
}{
	{
		"when given a string",
		"walking",
		[]string{"walking"},
	},

	{
		"when given nil",
		nil,
		[]string{},
	},

	{
		"when given int",
		1,
		[]string{},
	},

	{
		"when given empty string",
		"",
		[]string{""},
	},

	{
		"when given a struct with a string",
		struct {
			string
		}{"walking"},
		[]string{"walking"},
	},

	{
		"when given a struct with 2 strings",
		struct {
			name     string
			location string
		}{"Chris", "Agra"},
		[]string{"Chris", "Agra"},
	},

	{
		"when given a struct with a string and a non string",
		struct {
			name string
			age  int
		}{"Chris", 20},
		[]string{"Chris"},
	},

	{
		"when given a nested struct",
		struct {
			name    string
			profile profile
		}{"Chris", profile{20, "Agra"}},
		[]string{"Chris", "Agra"},
	},

	{
		"when given pointers to things",
		&struct {
			name    string
			profile profile
		}{"Chris", profile{20, "Agra"}},
		[]string{"Chris", "Agra"},
	},
}

func TestWalk(t *testing.T) {
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			got := make([]string, 0)
			fn := func(str string) { got = append(got, str) }

			reflection.Walk(test.input, fn)

			AssertEqual(t, got, test.want)
		})
	}
}