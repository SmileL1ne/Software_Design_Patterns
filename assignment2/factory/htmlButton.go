package main

import (
	"fmt"
)

type HtmlButton struct {
}

func (h *HtmlButton) render() {
	fmt.Println("*Http Button*")
	h.onClick()
}

func (h *HtmlButton) onClick() {
	fmt.Println("Click! It says: 'Hello Mustik!'")
}
