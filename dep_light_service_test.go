package gobenchreceiver

import "testing"

func Benchmark_DepServiceLightStruct_UpStruct(b *testing.B) {
	up := BasicServiceStruct{
		StringField: "some-long-string-here",
		IntField:    42,
		FloatField:  42.42,
	}
	s := DepServiceLightStruct{
		A: up,
	}
	var r int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = s.SomeMethod(100)
	}
	result = r
}

func Benchmark_DepServiceLightStruct_UpPointer(b *testing.B) {
	up := BasicServicePointer{
		StringField: "some-long-string-here",
		IntField:    42,
		FloatField:  42.42,
	}
	s := DepServiceLightStruct{
		A: &up,
	}
	var r int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = s.SomeMethod(100)
	}
	result = r
}

func Benchmark_DepServiceLightPointer_UpStruct(b *testing.B) {
	up := BasicServiceStruct{
		StringField: "some-long-string-here",
		IntField:    42,
		FloatField:  42.42,
	}
	s := &DepServiceLightPointer{
		A: up,
	}
	var r int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = s.SomeMethod(100)
	}
	result = r
}

func Benchmark_DepServiceLightPointer_UpPointer(b *testing.B) {
	up := BasicServicePointer{
		StringField: "some-long-string-here",
		IntField:    42,
		FloatField:  42.42,
	}
	s := &DepServiceLightPointer{
		A: &up,
	}
	var r int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = s.SomeMethod(100)
	}
	result = r
}

func Benchmark_DepServiceLightStruct_UpStructMany_10(b *testing.B) {
	up := BasicServicePointer{
		StringField: "some-long-string-here",
		IntField:    42,
		FloatField:  42.42,
	}

	l0 := DepServiceLightStruct{A: &up}
	l1 := DepServiceLightStruct{A: l0}
	l2 := DepServiceLightStruct{A: l1}
	l3 := DepServiceLightStruct{A: l2}
	l4 := DepServiceLightStruct{A: l3}
	l5 := DepServiceLightStruct{A: l4}
	l6 := DepServiceLightStruct{A: l5}
	l7 := DepServiceLightStruct{A: l6}
	l8 := DepServiceLightStruct{A: l7}
	l9 := DepServiceLightStruct{A: l8}

	s := DepServiceLightStruct{
		A: l9,
	}
	var r int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = s.SomeMethod(100)
	}
	result = r
}

func Benchmark_DepServiceLightPointer_UpPointerMany_10(b *testing.B) {
	up := BasicServicePointer{
		StringField: "some-long-string-here",
		IntField:    42,
		FloatField:  42.42,
	}

	l0 := &DepServiceLightPointer{A: &up}
	l1 := &DepServiceLightPointer{A: l0}
	l2 := &DepServiceLightPointer{A: l1}
	l3 := &DepServiceLightPointer{A: l2}
	l4 := &DepServiceLightPointer{A: l3}
	l5 := &DepServiceLightPointer{A: l4}
	l6 := &DepServiceLightPointer{A: l5}
	l7 := &DepServiceLightPointer{A: l6}
	l8 := &DepServiceLightPointer{A: l7}
	l9 := &DepServiceLightPointer{A: l8}

	s := &DepServiceLightPointer{
		A: l9,
	}
	var r int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = s.SomeMethod(100)
	}
	result = r
}

func Benchmark_DepServiceLightStruct_UpStructMany_50(b *testing.B) {
	up := BasicServicePointer{
		StringField: "some-long-string-here",
		IntField:    42,
		FloatField:  42.42,
	}

	l0 := DepServiceLightStruct{A: &up}
	l1 := DepServiceLightStruct{A: l0}
	l2 := DepServiceLightStruct{A: l1}
	l3 := DepServiceLightStruct{A: l2}
	l4 := DepServiceLightStruct{A: l3}
	l5 := DepServiceLightStruct{A: l4}
	l6 := DepServiceLightStruct{A: l5}
	l7 := DepServiceLightStruct{A: l6}
	l8 := DepServiceLightStruct{A: l7}
	l9 := DepServiceLightStruct{A: l8}
	l10 := DepServiceLightStruct{A: l9}
	l11 := DepServiceLightStruct{A: l10}
	l12 := DepServiceLightStruct{A: l11}
	l13 := DepServiceLightStruct{A: l12}
	l14 := DepServiceLightStruct{A: l13}
	l15 := DepServiceLightStruct{A: l14}
	l16 := DepServiceLightStruct{A: l15}
	l17 := DepServiceLightStruct{A: l16}
	l18 := DepServiceLightStruct{A: l17}
	l19 := DepServiceLightStruct{A: l18}
	l20 := DepServiceLightStruct{A: l19}
	l21 := DepServiceLightStruct{A: l20}
	l22 := DepServiceLightStruct{A: l21}
	l23 := DepServiceLightStruct{A: l22}
	l24 := DepServiceLightStruct{A: l23}
	l25 := DepServiceLightStruct{A: l24}
	l26 := DepServiceLightStruct{A: l25}
	l27 := DepServiceLightStruct{A: l26}
	l28 := DepServiceLightStruct{A: l27}
	l29 := DepServiceLightStruct{A: l28}
	l30 := DepServiceLightStruct{A: l29}
	l31 := DepServiceLightStruct{A: l30}
	l32 := DepServiceLightStruct{A: l31}
	l33 := DepServiceLightStruct{A: l32}
	l34 := DepServiceLightStruct{A: l33}
	l35 := DepServiceLightStruct{A: l34}
	l36 := DepServiceLightStruct{A: l35}
	l37 := DepServiceLightStruct{A: l36}
	l38 := DepServiceLightStruct{A: l37}
	l39 := DepServiceLightStruct{A: l38}
	l40 := DepServiceLightStruct{A: l39}
	l41 := DepServiceLightStruct{A: l40}
	l42 := DepServiceLightStruct{A: l41}
	l43 := DepServiceLightStruct{A: l42}
	l44 := DepServiceLightStruct{A: l43}
	l45 := DepServiceLightStruct{A: l44}
	l46 := DepServiceLightStruct{A: l45}
	l47 := DepServiceLightStruct{A: l46}
	l48 := DepServiceLightStruct{A: l47}
	l49 := DepServiceLightStruct{A: l48}

	s := DepServiceLightStruct{
		A: l49,
	}
	var r int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = s.SomeMethod(100)
	}
	result = r
}

func Benchmark_DepServiceLightPointer_UpPointerMany_50(b *testing.B) {
	up := BasicServicePointer{
		StringField: "some-long-string-here",
		IntField:    42,
		FloatField:  42.42,
	}

	l0 := &DepServiceLightPointer{A: &up}
	l1 := &DepServiceLightPointer{A: l0}
	l2 := &DepServiceLightPointer{A: l1}
	l3 := &DepServiceLightPointer{A: l2}
	l4 := &DepServiceLightPointer{A: l3}
	l5 := &DepServiceLightPointer{A: l4}
	l6 := &DepServiceLightPointer{A: l5}
	l7 := &DepServiceLightPointer{A: l6}
	l8 := &DepServiceLightPointer{A: l7}
	l9 := &DepServiceLightPointer{A: l8}
	l10 := &DepServiceLightPointer{A: l9}
	l11 := &DepServiceLightPointer{A: l10}
	l12 := &DepServiceLightPointer{A: l11}
	l13 := &DepServiceLightPointer{A: l12}
	l14 := &DepServiceLightPointer{A: l13}
	l15 := &DepServiceLightPointer{A: l14}
	l16 := &DepServiceLightPointer{A: l15}
	l17 := &DepServiceLightPointer{A: l16}
	l18 := &DepServiceLightPointer{A: l17}
	l19 := &DepServiceLightPointer{A: l18}
	l20 := &DepServiceLightPointer{A: l19}
	l21 := &DepServiceLightPointer{A: l20}
	l22 := &DepServiceLightPointer{A: l21}
	l23 := &DepServiceLightPointer{A: l22}
	l24 := &DepServiceLightPointer{A: l23}
	l25 := &DepServiceLightPointer{A: l24}
	l26 := &DepServiceLightPointer{A: l25}
	l27 := &DepServiceLightPointer{A: l26}
	l28 := &DepServiceLightPointer{A: l27}
	l29 := &DepServiceLightPointer{A: l28}
	l30 := &DepServiceLightPointer{A: l29}
	l31 := &DepServiceLightPointer{A: l30}
	l32 := &DepServiceLightPointer{A: l31}
	l33 := &DepServiceLightPointer{A: l32}
	l34 := &DepServiceLightPointer{A: l33}
	l35 := &DepServiceLightPointer{A: l34}
	l36 := &DepServiceLightPointer{A: l35}
	l37 := &DepServiceLightPointer{A: l36}
	l38 := &DepServiceLightPointer{A: l37}
	l39 := &DepServiceLightPointer{A: l38}
	l40 := &DepServiceLightPointer{A: l39}
	l41 := &DepServiceLightPointer{A: l40}
	l42 := &DepServiceLightPointer{A: l41}
	l43 := &DepServiceLightPointer{A: l42}
	l44 := &DepServiceLightPointer{A: l43}
	l45 := &DepServiceLightPointer{A: l44}
	l46 := &DepServiceLightPointer{A: l45}
	l47 := &DepServiceLightPointer{A: l46}
	l48 := &DepServiceLightPointer{A: l47}
	l49 := &DepServiceLightPointer{A: l48}

	s := &DepServiceLightPointer{
		A: l49,
	}
	var r int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = s.SomeMethod(100)
	}
	result = r
}
