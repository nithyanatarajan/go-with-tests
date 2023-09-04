package concurrency

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	durationA := measureResponseTime(a)
	durationB := measureResponseTime(b)

	if durationA < durationB {
		return a
	}
	return b
}

func RacerWithSelect(a, b string,
	timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		checkResponse(url)
		defer close(ch)
	}()
	return ch
}

func measureResponseTime(url string) time.Duration {
	startTime := time.Now()
	checkResponse(url)
	return time.Since(startTime)
}

func checkResponse(url string) {
	_, _ = http.Get(url)
}
