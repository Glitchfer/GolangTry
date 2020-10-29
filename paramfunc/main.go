package main

import (
	"fmt"
	"strings"
)

func main() {
	data := []string{"beybi", "ayi", "miya"}
	arrData := filter(data, func(setData string) bool {
		return strings.Contains(setData, "a")
	})

	fmt.Println(arrData)
}

func filter(data []string, callback func(string) bool) []string {
	var result []string

	for _, value := range data {

		if filtered := callback(value); filtered {
			result = append(result, value)
		}
	}

	return result
}
