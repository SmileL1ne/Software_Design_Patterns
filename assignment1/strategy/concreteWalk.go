package main

type Walk struct {
}

func (w *Walk) GetRoute(start, end string) string {
	return "Your walk route here: " + start + " -> " + end
}
