package main

import "fmt"

func main() {
	breaK()
	continuE()
}

func breaK() {

	for i := 0; i < 10; i++ {

		if i == 5 {
			fmt.Println("loop ke ", i, " Break here")
			break
		} else {
			fmt.Println("loop ke ", i)
		}
	}
}

func continuE() {

	for i := 0; i < 10; i++ {

		if i%2 == 0 {
			fmt.Println("Loncat")
			continue
		} else {
			fmt.Println("Perulangan ke ", i)
		}

	}

}
