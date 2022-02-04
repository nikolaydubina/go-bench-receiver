package gobenchreceiver

import "testing"

func Benchmark_EmptyServiceStruct(b *testing.B) {
	s := EmptyServiceStruct{}
	var r int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = s.SomeMethod(1000)
	}
	result = r
}

func Benchmark_EmptyServicePointer(b *testing.B) {
	s := &EmptyServicePointer{}
	var r int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = s.SomeMethod(1000)
	}
	result = r
}
