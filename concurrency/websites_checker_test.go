package concurrency_test

import (
	"github.com/nithyanatarajan/go-with-tests/concurrency"
	. "github.com/nithyanatarajan/go-with-tests/test"
	"testing"
)

func mockWebsiteResponses(url string) bool {
	if url == "https://allow.com" {
		return true
	}
	return false
}

func TestCheckWebsites(t *testing.T) {
	urls := []string{"https://allow.com", "https://a.com", "https://b.com"}
	want := map[string]bool{"https://allow.com": true, "https://a.com": false, "https://b.com": false}

	t.Run("should return response for given website", func(t *testing.T) {
		got := concurrency.CheckWebsites(mockWebsiteResponses, urls)

		AssertEqual(t, got, want)
	})
}
