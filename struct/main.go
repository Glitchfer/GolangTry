package main

import "fmt"

// struct adalah cara digolang untuk membuat object
type student struct {
	name   string
	class  string
	number int
}

func main() {
	// normal struct
	var item student

	item.name = "bambang"
	item.class = "2"

	fmt.Println(item)
	fmt.Println(item.number)

	fmt.Println("=========End of normal struct===========")
	// ===================================================
	// struct with object
	data := student{
		name:   "Bimbap",
		class:  "10",
		number: 31,
	}

	fmt.Println(data)
	fmt.Println("=========End of struct with object===========")
	// ====================================================
	// pointer struct
	var point *student

	point = &student{
		name: "obama",
	}

	fmt.Println(point)

	fmt.Println("==========End of pointer struct=============")
}
