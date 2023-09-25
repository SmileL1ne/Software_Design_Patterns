package main

func main() {
	phone := newItem("Phone")

	customerFirst := &Customer{id: "myEmail@gmail.com"}
	customerSecond := &Customer{id: "notMyEmail@gmail.com"}

	phone.register(customerFirst)
	phone.register(customerSecond)

	phone.updateAvailability()
}
