package print_test

import (
	"bytes"
	"testing"

	"github.com/nithyanatarajan/go-with-tests/print"
	. "github.com/nithyanatarajan/go-with-tests/test"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	t.Run("should count down from 3", func(t *testing.T) {
		buffer := bytes.Buffer{}
		sleeper := SpySleeper{}

		print.Countdown(&buffer, &sleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		AssertEqual(t, got, want)
		AssertEqual(t, sleeper.Calls, 3)
	})
}
