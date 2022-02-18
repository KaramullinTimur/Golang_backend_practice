package api

type API interface {
	HandleHttpRequest()
	getItems()
	addItem()
}
