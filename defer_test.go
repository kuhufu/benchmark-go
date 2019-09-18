package benchmark_go

import (
	"sync"
	"testing"
)

/*
   goos: windows
   goarch: amd64
   BenchmarkDeferLock-4   	30000000	        46.6 ns/op
   BenchmarkLock-4        	100000000	        15.4 ns/op
   PASS
   go version go1.12.7 windows/amd64
*/

func BenchmarkDeferLock(b *testing.B) {
	mu := sync.Mutex{}
	f := func() {
		mu.Lock()
		defer mu.Unlock()
	}
	for i := 0; i < b.N; i++ {
		f()
	}
}

func BenchmarkLock(b *testing.B) {
	mu := sync.Mutex{}
	f := func() {
		mu.Lock()
		mu.Unlock()
	}
	for i := 0; i < b.N; i++ {
		f()
	}
}
