package main

import "fmt"

var dialog Dialog

func main() {
	configure()
	runBusinessLogic()
}

func configure() {
	var choice string

	fmt.Println("Which dialog window do you want?(html, go): ")

	for {
		fmt.Scan(&choice)
		switch choice {
		case "html":
			dialog = &HtmlDialog{}
			return
		case "go":
			dialog = &GoDialog{}
			return
		default:
			fmt.Println("Try again...")
		}
	}
}

func runBusinessLogic() {
	okButton := dialog.createButton()
	okButton.render()
}
