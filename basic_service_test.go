package gobenchreceiver

import "testing"

var result int

func Benchmark_BasicServiceStruct(b *testing.B) {
	s := BasicServiceStruct{
		StringField: "some-long-string-here",
		IntField:    42,
		FloatField:  42.42,
	}
	var r int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = s.SomeMethod(1000)
	}
	result = r
}

func Benchmark_BasicServicePointer(b *testing.B) {
	s := &BasicServicePointer{
		StringField: "some-long-string-here",
		IntField:    42,
		FloatField:  42.42,
	}
	var r int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = s.SomeMethod(1000)
	}
	result = r
}
