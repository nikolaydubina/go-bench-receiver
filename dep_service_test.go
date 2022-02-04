package gobenchreceiver

import "testing"

func Benchmark_DepServiceStruct_UpStruct(b *testing.B) {
	up := BasicServiceStruct{
		StringField: "some-long-string-here",
		IntField:    42,
		FloatField:  42.42,
	}
	s := DepServiceStruct{
		A:  up,
		U0: 1, U2: 2, U3: 3, U4: 4, U5: 5, U6: 6, U7: 7, U8: 8, U9: 9,
	}
	var r int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = s.SomeMethod(100)
	}
	result = r
}

func Benchmark_DepServiceStruct_UpPointer(b *testing.B) {
	up := BasicServicePointer{
		StringField: "some-long-string-here",
		IntField:    42,
		FloatField:  42.42,
	}
	s := DepServiceStruct{
		A:  &up,
		U0: 1, U2: 2, U3: 3, U4: 4, U5: 5, U6: 6, U7: 7, U8: 8, U9: 9,
	}
	var r int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = s.SomeMethod(100)
	}
	result = r
}

func Benchmark_DepServicePointer_UpStruct(b *testing.B) {
	up := BasicServiceStruct{
		StringField: "some-long-string-here",
		IntField:    42,
		FloatField:  42.42,
	}
	s := &DepServicePointer{
		A:  up,
		U0: 1, U2: 2, U3: 3, U4: 4, U5: 5, U6: 6, U7: 7, U8: 8, U9: 9,
	}
	var r int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = s.SomeMethod(100)
	}
	result = r
}

func Benchmark_DepServicePointer_UpPointer(b *testing.B) {
	up := BasicServicePointer{
		StringField: "some-long-string-here",
		IntField:    42,
		FloatField:  42.42,
	}
	s := &DepServicePointer{
		A:  &up,
		U0: 1, U2: 2, U3: 3, U4: 4, U5: 5, U6: 6, U7: 7, U8: 8, U9: 9,
	}
	var r int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = s.SomeMethod(100)
	}
	result = r
}

func Benchmark_DepServiceStruct_UpStructMany_10(b *testing.B) {
	up := BasicServicePointer{
		StringField: "some-long-string-here",
		IntField:    42,
		FloatField:  42.42,
	}

	l0 := DepServiceStruct{A: &up, U0: 1, U2: 2, U3: 3, U4: 4, U5: 5, U6: 6, U7: 7, U8: 8, U9: 9}
	l1 := DepServiceStruct{A: l0, U0: 1, U2: 2, U3: 3, U4: 4, U5: 5, U6: 6, U7: 7, U8: 8, U9: 9}
	l2 := DepServiceStruct{A: l1, U0: 1, U2: 2, U3: 3, U4: 4, U5: 5, U6: 6, U7: 7, U8: 8, U9: 9}
	l3 := DepServiceStruct{A: l2, U0: 1, U2: 2, U3: 3, U4: 4, U5: 5, U6: 6, U7: 7, U8: 8, U9: 9}
	l4 := DepServiceStruct{A: l3, U0: 1, U2: 2, U3: 3, U4: 4, U5: 5, U6: 6, U7: 7, U8: 8, U9: 9}
	l5 := DepServiceStruct{A: l4, U0: 1, U2: 2, U3: 3, U4: 4, U5: 5, U6: 6, U7: 7, U8: 8, U9: 9}
	l6 := DepServiceStruct{A: l5, U0: 1, U2: 2, U3: 3, U4: 4, U5: 5, U6: 6, U7: 7, U8: 8, U9: 9}
	l7 := DepServiceStruct{A: l6, U0: 1, U2: 2, U3: 3, U4: 4, U5: 5, U6: 6, U7: 7, U8: 8, U9: 9}
	l8 := DepServiceStruct{A: l7, U0: 1, U2: 2, U3: 3, U4: 4, U5: 5, U6: 6, U7: 7, U8: 8, U9: 9}
	l9 := DepServiceStruct{A: l8, U0: 1, U2: 2, U3: 3, U4: 4, U5: 5, U6: 6, U7: 7, U8: 8, U9: 9}

	s := DepServiceStruct{
		A: l9,
	}
	var r int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = s.SomeMethod(100)
	}
	result = r
}

func Benchmark_DepServicePointer_UpPointerMany_10(b *testing.B) {
	up := BasicServicePointer{
		StringField: "some-long-string-here",
		IntField:    42,
		FloatField:  42.42,
	}

	l0 := &DepServicePointer{A: &up, U0: 1, U2: 2, U3: 3, U4: 4, U5: 5, U6: 6, U7: 7, U8: 8, U9: 9}
	l1 := &DepServicePointer{A: l0, U0: 1, U2: 2, U3: 3, U4: 4, U5: 5, U6: 6, U7: 7, U8: 8, U9: 9}
	l2 := &DepServicePointer{A: l1, U0: 1, U2: 2, U3: 3, U4: 4, U5: 5, U6: 6, U7: 7, U8: 8, U9: 9}
	l3 := &DepServicePointer{A: l2, U0: 1, U2: 2, U3: 3, U4: 4, U5: 5, U6: 6, U7: 7, U8: 8, U9: 9}
	l4 := &DepServicePointer{A: l3, U0: 1, U2: 2, U3: 3, U4: 4, U5: 5, U6: 6, U7: 7, U8: 8, U9: 9}
	l5 := &DepServicePointer{A: l4, U0: 1, U2: 2, U3: 3, U4: 4, U5: 5, U6: 6, U7: 7, U8: 8, U9: 9}
	l6 := &DepServicePointer{A: l5, U0: 1, U2: 2, U3: 3, U4: 4, U5: 5, U6: 6, U7: 7, U8: 8, U9: 9}
	l7 := &DepServicePointer{A: l6, U0: 1, U2: 2, U3: 3, U4: 4, U5: 5, U6: 6, U7: 7, U8: 8, U9: 9}
	l8 := &DepServicePointer{A: l7, U0: 1, U2: 2, U3: 3, U4: 4, U5: 5, U6: 6, U7: 7, U8: 8, U9: 9}
	l9 := &DepServicePointer{A: l8, U0: 1, U2: 2, U3: 3, U4: 4, U5: 5, U6: 6, U7: 7, U8: 8, U9: 9}

	s := &DepServicePointer{
		A: l9,
	}
	var r int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = s.SomeMethod(100)
	}
	result = r
}
