package main

import "fmt"

func funcWithReturnValue(name string) string {
	return "Hello " + name
}

func main() {

	result := funcWithReturnValue("Arif")
	fmt.Println(result)
}
