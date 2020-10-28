package main

import "fmt"

func main() {
	fmt.Println("================(js while)=================")
	// for sperti while
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	fmt.Println("================(normal for)=================")

	// for dengan statement (seperti di js tapi ga pake ())
	for counter := 0; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	fmt.Println("================( How to extract arr with for )=================")

	arr := []string{"Eko", "Kurniawan", "Khennedy"}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	fmt.Println("================(FOR Range)=================")

	// untuk mengambil semua data (value & index) yang ada di array
	// seperti foreach & map di javascript

	character := []string{"Templar Assasin", "Invoker", "Litch", "Puck", "Dokter Witch"}

	for index, value := range character {
		fmt.Println("index ", index, " = ", value)
	}

	fmt.Println("==========")

	// jika hanya salah satu variabel yg digunakan maka gantikan "_" di variabel tsb
	tipe := []string{"Strength", "Agility", "Intelligent"}

	for _, value := range tipe {
		fmt.Println("Tipe ", value)
	}

	fmt.Println("=== for range untuk map ===")

	hero := make(map[string]string)
	hero["nama"] = "Saitama"
	hero["hobi"] = "Main ps"

	for key, value := range hero {
		fmt.Println(key, "=", value)
	}
}
