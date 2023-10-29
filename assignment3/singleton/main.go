package main

import "fmt"

func main() {
	instance1 := getIntance()
	instance1.name = "neMus"
	instance1.age = 90
	instance2 := getIntance()
	instance2.name = "Mus"
	instance2.age = 20

	for i := 0; i < 50; i++ {
		go getIntance()
	}

	if instance1 == instance2 {
		fmt.Println("Yay! Singleton worked!")
		fmt.Println("Stuct from instance1:", instance1)
		fmt.Println("Stuct from instance2:", instance2)
	} else {
		fmt.Println("Nope, singleton didn't work")
	}
}
