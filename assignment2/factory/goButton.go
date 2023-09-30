package main

import (
	"fmt"
	"os"
)

var count int

type GoButton struct {
}

func (g *GoButton) render() {
	fmt.Println("I couldn't make button with gtk gui library cuz it's too complex so here is my version of go button: ")
	fmt.Println("*Button*")
	fmt.Println("*Click me by typing 'Click'")
	var click string
	for {
		_, err := fmt.Scan(&click)
		if err != nil {
			panic(err)
		}
		if click == "Click" {
			g.onClick()
		} else {
			fmt.Println("Try again...")
			count++
		}
	}
}

func (g *GoButton) onClick() {
	if count > 0 {
		fmt.Println("Finally...\nBye... guess...")
		count = 0
	} else {
		fmt.Println("Bye!")
	}
	os.Exit(0)
}
