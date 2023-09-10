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
		"when given empty struct",
		struct{}{},
		[]string{},
	},

	{
		"when given pointers to things",
		&struct {
			name    string
			profile profile
		}{"Chris", profile{20, "Agra"}},
		[]string{"Chris", "Agra"},
	},

	{
		"when given string slice",
		[]string{"Chris", "Agra"},
		[]string{"Chris", "Agra"},
	},

	{
		"when given int slice",
		[]int{10, 20},
		[]string{},
	},

	{
		"when given struct slices",
		[]profile{
			{20, "Agra"},
			{33, "London"},
		},
		[]string{"Agra", "London"},
	},

	{
		"when given empty slice",
		[]profile{},
		[]string{},
	},

	{
		"when given string array",
		[2]string{"Chris", "Agra"},
		[]string{"Chris", "Agra"},
	},

	{
		"when given empty array",
		[1]profile{},
		[]string{""},
	},

	{
		"when given empty map",
		map[string]string{},
		[]string{},
	},

	{
		"when given function",
		func() []profile {
			return []profile{
				{20, "Agra"},
				{33, "London"},
			}
		},
		[]string{"Agra", "London"},
	},

	{
		"when given function with args",
		func(pro profile) []profile {
			return []profile{pro}
		}(profile{20, "Agra"}),
		[]string{"Agra"},
	},

	{
		"when given empty function",
		func() {},
		[]string{},
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

	t.Run("when given map", func(t *testing.T) {
		input := map[string]string{"name": "Chris", "location": "Agra"}
		got := make([]string, 0)
		fn := func(str string) { got = append(got, str) }

		reflection.Walk(input, fn)

		AssertEqual(t, len(got), 2)
		AssertContains(t, got, "Chris")
		AssertContains(t, got, "Agra")
	})

	t.Run("when given mixed type", func(t *testing.T) {
		input := &struct {
			name      string
			profiles  [1]profile
			relations map[string][]string
		}{"Chris",
			[1]profile{{20, "Agra"}},
			map[string][]string{
				"friends": {"Polly", "Molly"},
				"enemies": {"Jolly"}},
		}
		got := make([]string, 0)
		fn := func(str string) { got = append(got, str) }

		reflection.Walk(input, fn)

		AssertEqual(t, len(got), 5)
		AssertContains(t, got, "Chris")
		AssertContains(t, got, "Agra")
		AssertContains(t, got, "Polly")
		AssertContains(t, got, "Molly")
		AssertContains(t, got, "Jolly")
	})
}
