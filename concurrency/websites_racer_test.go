package concurrency_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/nithyanatarajan/go-with-tests/concurrency"
	. "github.com/nithyanatarajan/go-with-tests/test"
)

func TestRacer(t *testing.T) {
	t.Run("should return fast responding website",
		func(t *testing.T) {
			slowServer := Server(func() {
				time.Sleep(20 * time.Millisecond)
			})
			fastServer := Server(func() {})
			defer slowServer.Close()
			defer fastServer.Close()

			slowURL := slowServer.URL
			fastURL := fastServer.URL
			want := fastURL

			got := concurrency.Racer(slowURL, fastURL)

			AssertEqual(t, got, want)
		})
}

func TestRacerWithSelect(t *testing.T) {
	oneSecondTimeout := 1 * time.Second
	t.Run("should return fast responding website",
		func(t *testing.T) {
			slowServer := Server(func() {
				time.Sleep(20 * time.Millisecond)
			})
			fastServer := Server(func() {})
			defer slowServer.Close()
			defer fastServer.Close()

			slowURL := slowServer.URL
			fastURL := fastServer.URL
			want := fastURL

			got, err := concurrency.RacerWithSelect(slowURL, fastURL,
				oneSecondTimeout)

			AssertEqual(t, got, want)
			AssertNoError(t, err)
		})

	t.Run("should return error if response time is more than 1 seconds",
		func(t *testing.T) {
			serverA := Server(func() {
				time.Sleep(3 * time.Second)
			})
			serverB := Server(func() {
				time.Sleep(2 * time.Second)
			})

			defer serverA.Close()
			defer serverB.Close()

			got, err := concurrency.RacerWithSelect(
				serverA.URL,
				serverB.URL,
				oneSecondTimeout)

			AssertError(t, err)
			AssertEqual(t, got, "")
		})

	t.Run("should return not error "+
		"if response time is more than 1 seconds only for one server",
		func(t *testing.T) {
			serverA := Server(func() {
				time.Sleep(900 * time.Millisecond)
			})
			serverB := Server(func() {
				time.Sleep(oneSecondTimeout)
			})

			defer serverA.Close()
			defer serverB.Close()
			want := serverA.URL

			got, err := concurrency.RacerWithSelect(
				serverA.URL,
				serverB.URL,
				oneSecondTimeout)

			AssertEqual(t, got, want)
			AssertNoError(t, err)
		})
}

func Server(action func()) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			action()
			w.WriteHeader(http.StatusOK)
		}))
}
