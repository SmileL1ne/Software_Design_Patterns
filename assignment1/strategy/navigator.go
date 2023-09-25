package main

type Navigator interface {
	GetRoute(start, end string) string
}
