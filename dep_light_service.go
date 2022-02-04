package gobenchreceiver

type DepServiceLightStruct struct {
	A DepService
}

func (s DepServiceLightStruct) SomeMethod(a int) int {
	return s.A.SomeMethod(a) + 1
}

type DepServiceLightPointer struct {
	A DepService
}

func (s *DepServiceLightPointer) SomeMethod(a int) int {
	if s == nil {
		return 0
	}
	return s.A.SomeMethod(a) + 1
}
