package main

var navigator Navigator

type Context struct {
}

func (c *Context) SetNavigator(str Navigator) {
	navigator = str
}

func (c *Context) GetRoute(start, end string) string {
	return navigator.GetRoute(start, end)
}
