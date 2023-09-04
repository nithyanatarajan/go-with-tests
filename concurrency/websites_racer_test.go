package concurrency_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/nithyanatarajan/go-with-tests/concurrency"
	"github.com/nithyanatarajan/go-with-tests/test"
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

			test.AssertEqual(t, got, want)
		})
}

func Server(action func()) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			action()
			w.WriteHeader(http.StatusOK)
		}))
}
