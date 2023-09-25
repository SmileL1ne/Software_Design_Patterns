package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	context := Context{}
	var start string
	var end string
	var action string

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your starting point: ")
	start, _ = reader.ReadString('\n')
	start = strings.TrimSpace(start)
	fmt.Print("Enter your destination: ")
	end, _ = reader.ReadString('\n')
	end = strings.TrimSpace(end)
	fmt.Print("Enter desired action: ")
	action, _ = reader.ReadString('\n')
	action = strings.TrimSpace(action)

	switch action {
	case "walk":
		context.SetNavigator(&Walk{})
	case "bike":
		context.SetNavigator(&Bike{})
	case "car":
		context.SetNavigator(&Car{})
	default:
		fmt.Println("Invalid route method")
	}

	result := context.GetRoute(start, end)

	fmt.Println(result)

}
