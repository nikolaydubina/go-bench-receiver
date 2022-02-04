package gobenchreceiver

type BasicServiceStruct struct {
	StringField string
	IntField    int
	FloatField  float64
}

func (s BasicServiceStruct) SomeMethod(a int) int {
	var q int
	if len(s.StringField) > 0 {
		q++
	}
	e := len(s.StringField) + s.IntField + int(s.FloatField) + q
	x := 0
	for i := 0; i < e; i++ {
		x = x * x % 1001
	}
	return x
}

type BasicServicePointer struct {
	StringField string
	IntField    int
	FloatField  float64
}

func (s *BasicServicePointer) SomeMethod(a int) int {
	if s == nil {
		return 0
	}
	var q int
	if len(s.StringField) > 0 {
		q++
	}
	e := len(s.StringField) + s.IntField + int(s.FloatField) + q
	x := 0
	for i := 0; i < e; i++ {
		x = x * x % 1001
	}
	return x
}
