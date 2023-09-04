package concurrency_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/nithyanatarajan/go-with-tests/concurrency"
)

func slowStubWebsiteResponses(_ string) bool {
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
