package main

import "fmt"

func main() {
	name := "Arifa"

	if name == "Arif" {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	fmt.Println("=======================")

	if name == "Arif" {
		fmt.Println("Welcome back ", name)
	} else if name == "Arifa" {
		fmt.Println("Hallo Arifa")
	} else {
		fmt.Println("Kamu bukan Arif")
	}

	fmt.Println("=======================")

	length := len(name)
	if length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

	// pendeklarasian variabel juga bisa didalam if, jadi seperti berikut ini

	if length := len(name); length < 5 {
		fmt.Println("Nama terlalu pendek")
	} else {
		fmt.Println("Panjang nama anda adalah ", len(name))
	}
}
