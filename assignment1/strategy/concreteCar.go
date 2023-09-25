package main

type Car struct {
}

func (c *Car) GetRoute(start, end string) string {
	return "Youre car route: " + start + " -> " + end
}
