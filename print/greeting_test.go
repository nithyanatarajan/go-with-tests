package print_test

import (
	"bytes"
	"testing"

	"github.com/nithyanatarajan/go-with-tests/print"
	. "github.com/nithyanatarajan/go-with-tests/test"
)

func TestGreet(t *testing.T) {
	t.Run("should Greet person", func(t *testing.T) {
		buffer := bytes.Buffer{}
		print.Greet(&buffer, "Chris")

		got := buffer.String()
		want := "Hello, Chris"
		AssertEqual(t, got, want)
	})
}
