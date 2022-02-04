package gobenchreceiver

type DepService interface {
	SomeMethod(a int) int
}

type DepServiceStruct struct {
	A  DepService
	U0 uint64
	U1 uint64
	U2 uint64
	U3 uint64
	U4 uint64
	U5 uint64
	U6 uint64
	U7 uint64
	U8 uint64
	U9 uint64
}

func (s DepServiceStruct) SomeMethod(a int) int {
	return s.A.SomeMethod(a) + 1
}

type DepServicePointer struct {
	A  DepService
	U0 uint64
	U1 uint64
	U2 uint64
	U3 uint64
	U4 uint64
	U5 uint64
	U6 uint64
	U7 uint64
	U8 uint64
	U9 uint64
}

func (s *DepServicePointer) SomeMethod(a int) int {
	if s == nil {
		return 0
	}
	return s.A.SomeMethod(a) + 1
}
