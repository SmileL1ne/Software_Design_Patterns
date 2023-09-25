package main

import "fmt"

type Item struct {
	observerList []Observer
	name         string
	inStock      bool
}

func newItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) updateAvailability() {
	fmt.Printf("Item %s now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}

func (i *Item) register(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *Item) derigester(o Observer) {
	i.observerList = removeFromSlice(i.observerList, o)
}

func (i *Item) notifyAll() {
	for _, o := range i.observerList {
		o.update(i.name)
	}
}

func removeFromSlice(list []Observer, o Observer) []Observer {
	for i, oi := range list {
		if oi.getID() == o.getID() {
			list[len(list)-1], list[i] = list[i], list[len(list)-1]
		}
		return list[:len(list)-1]
	}
	return list
}
