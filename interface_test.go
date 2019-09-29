package benchmark_go

import (
	"testing"
)

/*
给interface类型参数传递一个指针是一个好选择
goos: windows
goarch: amd64
BenchmarkPassInterface-4               	10000000	       244 ns/op
BenchmarkPassStruct-4                  	10000000	       224 ns/op
BenchmarkPassStructPtr-4               	2000000000	         0.28 ns/op
BenchmarkPassStructPtrToInterface-4    	2000000000	         0.28 ns/op
BenchmarkPassInterfaceToInterface-4    	2000000000	         0.28 ns/op
BenchmarkPassInterfaceToInterface2-4   	2000000000	         0.28 ns/op
PASS
*/

type Person struct {
	Name string
	Age  int64
	Data [128]int64
}

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Inter(&Person{}, 100)
	}
}

func Benchmark12(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Inter(Person{}, 10)
	}
}

func Benchmark21(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Inter2(Person{}, 10)
	}
}

func Inter(a interface{}, i int) {
	if i == 0 {
		return
	}

	Inter(a, i-1)
}

func Inter2(a Person, i int) {
	if i == 0 {
		return
	}

	Inter3(a)
	Inter(a, i-1)
}

func Inter3(a Person) {
	a.Age++
}

func funcInterfaceToStruct(a interface{}) {
	foo(a)
}

func funcStruct(a Person) {
	foo(a)
}

func funcStructPtr(a *Person) {
	foo(a)
}

func funcInterfaceToPtr(a interface{}) {
	foo(a)
}

func foo(a interface{}) {

}

func BenchmarkPassInterface(b *testing.B) {
	var p = Person{
		Name: "kuhufu",
		Age:  10,
	}
	for i := 0; i < b.N; i++ {
		funcInterfaceToStruct(p)
	}
}

func BenchmarkPassStruct(b *testing.B) {
	var p = Person{
		Name: "kuhufu",
		Age:  10,
	}
	for i := 0; i < b.N; i++ {
		funcStruct(p)
	}
}

func BenchmarkPassStructPtr(b *testing.B) {
	var p = Person{
		Name: "kuhufu",
		Age:  10,
	}
	for i := 0; i < b.N; i++ {
		funcStructPtr(&p)
	}
}

func BenchmarkPassStructPtrToInterface(b *testing.B) {
	var p = Person{
		Name: "kuhufu",
		Age:  10,
	}
	for i := 0; i < b.N; i++ {
		funcInterfaceToPtr(&p)
	}
}

func BenchmarkPassInterfaceToInterface(b *testing.B) {
	var p interface{} = Person{
		Name: "kuhufu",
		Age:  10,
	}
	for i := 0; i < b.N; i++ {
		funcInterfaceToStruct(p)
	}
}

func BenchmarkPassInterfaceToInterface2(b *testing.B) {
	var p interface{} = Person{
		Name: "kuhufu",
		Age:  10,
	}
	for i := 0; i < b.N; i++ {
		funcInterfaceToStruct(p)
	}
}
