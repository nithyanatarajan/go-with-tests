package concurrency

import (
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

func measureResponseTime(url string) time.Duration {
	startTime := time.Now()
	_, _ = http.Get(url)
	return time.Since(startTime)
}
