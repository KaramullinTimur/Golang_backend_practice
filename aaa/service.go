package aaa

// type Serivicer interface {
// 	ServiceFunc() int
// }

type service struct {
	m Mocker
}

func NewService(i int) service {
	return service{Cock1{2}}
}

func (s service) ServiceFunc() int {
	return 2 + s.m.mock()
}

type Mocker interface {
	mock() int
}

type Mock1 struct {
	num int
}

func (m Mock1) mock() int {
	return 1 + m.num
}

type Cock1 struct {
	num int
}

func (m Cock1) mock() int {
	return 11 * m.num
}
