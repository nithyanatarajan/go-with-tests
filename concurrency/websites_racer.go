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

func RacerCh(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
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
