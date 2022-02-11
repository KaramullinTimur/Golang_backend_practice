package aaa

func ApiFunc() int {
	srvc := NewService(3)

	a := srvc.ServiceFunc()
	return a + 2
}
