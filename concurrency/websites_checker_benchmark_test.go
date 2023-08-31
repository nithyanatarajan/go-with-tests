package concurrency_test

import (
	"fmt"
	"github.com/nithyanatarajan/go-with-tests/concurrency"
	"testing"
	"time"
)

func slowStubWebsiteResponses(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = fmt.Sprintf(`a{%d} url`, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		concurrency.CheckWebsites(slowStubWebsiteResponses, urls)
	}
}
