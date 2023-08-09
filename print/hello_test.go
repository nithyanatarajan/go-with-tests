package print_test

import (
	"testing"

	"github.com/nithyanatarajan/go-with-tests/print"
	. "github.com/nithyanatarajan/go-with-tests/test"
)

func TestHello(t *testing.T) {
	t.Run("should print Hello world when no name is given", func(t *testing.T) {
		const want = "Hello, World!"
		got := print.Hello("", "")
		AssertEqual(t, got, want)
	})

	t.Run("should print name if given", func(t *testing.T) {
		const want = "Hello, Chris!"
		got := print.Hello("Chris", "")
		AssertEqual(t, got, want)
	})

	t.Run("should print name in Spanish", func(t *testing.T) {
		const want = "Hola, Chris!"
		got := print.Hello("Chris", "Spanish")
		AssertEqual(t, got, want)
	})

	t.Run("should print name in French", func(t *testing.T) {
		const want = "Bonjour, Chris!"
		got := print.Hello("Chris", "French")
		AssertEqual(t, got, want)
	})
}
