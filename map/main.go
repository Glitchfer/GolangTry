package main

import (
	"fmt"
)

func main() {
	// membuat map
	hero := map[string]string{
		"name": "Phantom",
		"type": "Agility",
	}

	fmt.Println(hero)

	// mengubah/ menambahkan data di map dengan key
	hero["level"] = "1"
	fmt.Println(hero)

	// Menampilkan data di map dengan key
	fmt.Println(hero["name"])
	fmt.Println(hero["type"])

	// Melihat jumlah data di map
	fmt.Println(len(hero))

	// Membuat map baru
	status := make(map[string]int)
	status["HP"] = 30
	status["MP"] = 100
	status["Exp"] = 1350
	fmt.Println(status)

	// Menghapus data di map dengan key
	fmt.Println("=========================")
	delete(hero, "type")
	delete(status, "Exp")

	fmt.Println(hero)
	fmt.Println(status)
}
