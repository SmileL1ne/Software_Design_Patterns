package main

type Bike struct {
}

func (b *Bike) GetRoute(start, end string) string {
	return "Your bike route: " + start + " -> " + end
}
