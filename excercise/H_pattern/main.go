package main

import (
	"fmt"
)

func main() {
	segitiga()
}
func segitiga() {
	for x := 1; x <= 5; x++ {
		for y := 5; y >= x; y-- {
			if y == 5 || y == 1 {
				fmt.Print("*")
			} else {
				if x == 3 {
					fmt.Print("*")
				} else {
					fmt.Print("=")
				}
			}
		}
		for z := 1; z < x; z++ {
			if x == 3 {
				fmt.Print("*")
			} else {
				if (x == 2 && z == 1) || (x == 4 && z == 3) || (x == 5 && z == 4) {
					fmt.Print("*")
				} else {
					fmt.Print("=")
				}
			}
		}
		fmt.Println()
	}
}
