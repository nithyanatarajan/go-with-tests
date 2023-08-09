package print_test

import (
	"testing"

	"github.com/nithyanatarajan/go-with-tests/print"
	. "github.com/nithyanatarajan/go-with-tests/test"
)

func TestRepeat(t *testing.T) {
	type args struct {
		str   string
		times int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"should return empty if empty", args{}, ""},
		{
			"should repeat given string 5 times",
			args{str: "a", times: 5},
			"aaaaa",
		},
		{
			"should repeat given string 10 times",
			args{str: "ab", times: 10},
			"abababababababababab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := print.Repeat(tt.args.str, tt.args.times)
			AssertEqual(t, got, tt.want)
		})
	}
}
