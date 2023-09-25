package main

type Subject interface {
	register(obserber Observer)
	derigester(observer Observer)
	notifyAll()
}
