package main

import (
	"log"
)

const pi = 3.14

// reminder
// need example for using ponter on function

func main() {
	// printMessage("ini Hello wolrd dengan parameter")
	// printMessage2()
	// log.Println(getRange(10, 11))
	// log.Println(calculate(1.5))

	// log.Println(variadicCalculate(1, 1, 4, 2, 5, 7, 8, 9, 0, 12))
	var arrInt = []int{1, 21, 3, 4, 5, 33, 535, 353, 5, 353, 5, 35, 5353, 535, 35, 353, 535, 35, 35}
	resp, arrItem := findMax(arrInt, 2000)
	log.Println(resp)
	log.Println(arrItem())
}

// function tanpa parameter
// func printMessage2() {
// 	log.Println("ini Hello world tanpa parameter")
// }

// normal function
// function dengan parameter
// func printMessage(message string) { //bila menggunakan parameter harus diikuti dengan tipe datanya
// 	log.Println(message)
// }

// func dengan 1 return
// func getRange(min, max int) int {
// 	return max * min
// }

// func with 2 return
// func calculate(r float64) (float64, float64) {
// 	var area = pi * (r * r)

// 	var circle = 2 * pi * r

// 	return area, circle
// }

// func variadic
// func variadicCalculate(number ...int) float64 {
// 	var total int = 0
// 	for _, number := range number {
// 		total += number
// 	}

// 	var avg = float64(total) / float64(len(number))

// 	return avg
// }

// func as a return
func findMax(number []int, max int) (int, func() []int) {
	var res []int
	for _, p := range number {
		if p <= max {
			res = append(res, p)
		}
	}

	return len(res), func() []int {
		return res
	}
}
