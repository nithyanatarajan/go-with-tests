package print

import (
	"testing"
	"time"
)

const one = 1 * time.Second

func TestDefaultSleeper_Sleep(t *testing.T) {
	sleeper := &DefaultSleeper{}

	startTime := time.Now()
	sleeper.Sleep()
	actualDuration := time.Since(startTime).Round(one)

	expectedDuration := one
	if expectedDuration != actualDuration {
		t.Errorf("Sleep actualDuration is not 1 second. "+
			"Expected: %s, Actual: %s", expectedDuration, actualDuration)
	}
}
