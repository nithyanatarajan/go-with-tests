package print_test

import "testing"
import "github.com/nithyanatarajan/go-with-tests/print"

func TestHello(t *testing.T) {
	t.Run("should print Hello world when no name is given", func(t *testing.T) {
		const want = "Hello, World!"
		got := print.Hello("", "")
		assertEqual(t, got, want)

	})

	t.Run("should print name if given", func(t *testing.T) {
		const want = "Hello, Chris!"
		got := print.Hello("Chris", "")
		assertEqual(t, got, want)

	})

	t.Run("should print name in Spanish", func(t *testing.T) {
		const want = "Hola, Chris!"
		got := print.Hello("Chris", "Spanish")
		assertEqual(t, got, want)
	})

	t.Run("should print name in French", func(t *testing.T) {
		const want = "Bonjour, Chris!"
		got := print.Hello("Chris", "French")
		assertEqual(t, got, want)
	})
}

func assertEqual(t *testing.T, got, want string) {
	t.Helper() // When failed the line number reported will be in called function

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
