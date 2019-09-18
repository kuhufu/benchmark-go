package benchmark_go

import "testing"

/*
goos: windows
goarch: amd64
pkg: test/sliceCopy
Benchmark_slice_copy_builtin-4   	100000000	        15.1 ns/op
Benchmark_slice_copy_self-4      	30000000	        41.5 ns/op
PASS
*/

func Benchmark_slice_copy_builtin(b *testing.B) {
	dst := make([]int, 128)
	src := make([]int, 128)
	for i := 0; i < b.N; i++ {
		copy(dst, src)
	}
}

func Benchmark_slice_copy_self(b *testing.B) {
	dst := make([]int, 128)
	src := make([]int, 128)
	for i := 0; i < b.N; i++ {
		for i := 0; i < 128; i++ {
			dst[i] = src[i]
		}
	}
}