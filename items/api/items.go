package api

func (a *api) getItems() {
	// parse body
	// validate function
	// send status and message if error
	a.service.GetItems()
	// send status and message if error
}

func (a *api) addItem() {
	// parse body
	// validate function
	// send status and message if error
	a.service.AddItem()
	a.service.Items.AddItem()
	// send status and message if error
}
