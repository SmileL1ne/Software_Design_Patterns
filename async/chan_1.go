package main

import (
	"fmt"
	"runtime"
)

func main() {
	ch1 := make(chan int, 1)

	myFunc := func(in chan int) {
		val := <-in
		fmt.Println("GO: get from chan", val)
		fmt.Println("GO: after read from chan")
	}
	fmt.Println(runtime.NumGoroutine())
	go myFunc(ch1)
	fmt.Println("Second time call myFunc()")
	go myFunc(ch1)

	ch1 <- 42
	ch1 <- 100500

	fmt.Println(runtime.NumGoroutine())

	fmt.Println("MAIN: after put to chan")
	fmt.Scanln()
}
