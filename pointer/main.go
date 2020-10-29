package main

import "fmt"

func main() {
	var numberA int = 4
	// int pointer
	var numberB *int = &numberA

	fmt.Println(numberA)
	fmt.Println(&numberA)

	// get pointer value
	fmt.Println(*numberB)
	fmt.Println(numberB)

	// changePointerValue(numberB)
}

// func changePointerValue(item *int) {
// 	var numberA int = *item
// 	var numberB *int = &numberA

// 	fmt.Println(numberA)
// 	fmt.Println(&numberA)

// 	fmt.Println(*numberB)
// 	fmt.Println(numberB)

// 	numberA = 10

// 	fmt.Println(*numberB)
// 	fmt.Println(numberB)
// }
