package benchmark_go

import (
	"fmt"
	"testing"
	"time"
)

/*
降低时间的精度是一个加速的好选择
goos: windows
goarch: amd64
BenchmarkDate-4      	20000000	       110 ns/op
BenchmarkUnixSub-4   	2000000000	         0.35 ns/op
PASS
*/

func Test(t *testing.T) {
	now := time.Now()
	//fmt.Println(time.Duration((now.Unix() + 8*3600) % (3600 * 24) * 1e9))
	fmt.Println(now.Zone())
}

func BenchmarkDate(b *testing.B) {
	now := time.Now()
	for i := 0; i < b.N; i++ {
		_ = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Add(24 * time.Hour).Sub(now)
	}
}

func BenchmarkUnixSub(b *testing.B) {
	now := time.Now()
	for i := 0; i < b.N; i++ {
		_ = time.Duration((now.Unix() + 8*3600) % (3600 * 24)) * time.Second
	}
}
