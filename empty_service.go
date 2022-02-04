package gobenchreceiver

type EmptyServiceStruct struct{}

func (s EmptyServiceStruct) SomeMethod(a int) int {
	x := 0
	for i := 0; i < a; i++ {
		x = x * x % 1001
	}
	return x
}

type EmptyServicePointer struct{}

func (s *EmptyServicePointer) SomeMethod(a int) int {
	if s == nil {
		return 0
	}
	x := 0
	for i := 0; i < a; i++ {
		x = x * x % 1001
	}
	return x
}
