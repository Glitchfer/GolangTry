package main

import "fmt"

type customer struct {
	// penulisan normal
	// Name string
	// Address string
	// Age int

	// Bisa juga gini
	Name, Address string
	Age           int
}

// struct fucntion
func (customer customer) sayHai(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

// func (a customer) sayApaAja() {
// 	fmt.Println("wkwkwkwk lol", a.Name)
// }

func main() {
	// normal struct
	var data customer
	data.Name = "Arif"
	data.Address = "Indonesia"
	data.Age = 26

	fmt.Println(data)
	fmt.Println(data.Name)
	fmt.Println("========================")

	// struct literals
	data = customer{
		Name:    "Bambang",
		Address: "Cirebon",
		Age:     31,
	}

	fmt.Println(data)
	fmt.Println(data.Name)
	fmt.Println("========================")

	//  struct methods
	data.sayHai("Joko")
	// data.sayApaAja()
}
