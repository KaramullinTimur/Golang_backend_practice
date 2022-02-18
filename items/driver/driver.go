package driver

type Driver interface {
	GetItems()
	AddItem()
}

type Items interface {
	GetItems()
	AddItem()
}

type Status interface {
	SetStatus()
}
