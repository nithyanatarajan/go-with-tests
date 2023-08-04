package print_test

import "testing"
import "github.com/nithyanatarajan/go-with-tests/print"

func TestHello(t *testing.T) {
	t.Run("should print Hello world when no name is given", func(t *testing.T) {
		const want = "Hello, world!"
		if got := print.Hello(""); got != want {
			t.Errorf("Hello(\"\") = %v, want %v", got, want)
		}
	})

	t.Run("should print name if given", func(t *testing.T) {
		const want = "Hello, Chris!"
		if got := print.Hello("Chris"); got != want {
			t.Errorf("Hello(\"Chris\") = %v, want %v", got, want)
		}
	})
}
