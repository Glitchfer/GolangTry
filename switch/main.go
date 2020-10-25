package main

import "fmt"

func main() {
	// umur := 19

	// switch umur {
	// case 19:
	// 	fmt.Println("Bocah")
	// case 30:
	// 	fmt.Println("Dewasa")
	// default:
	// 	fmt.Println("Tua")
	// }

	// deklarasi variabel dapat diletakan didalam switch seperti ini
	switch umur := 19; umur < 19 {
	case true:
		fmt.Println("Umur anda belum mencukupi")
	case false:
		fmt.Println("Umur anda sudah mencukupi")
	}

	// switch dengan kondisi yang sebaris dengan case nya
	umur := 140
	switch {
	case umur < 19:
		fmt.Println("Anda masih bocah")
	case umur > 19 && umur <= 40:
		fmt.Println("Anda masih muda")
	case umur > 40 && umur <= 60:
		fmt.Println("Anda sudah tua")
	default:
		fmt.Println("Anda sudah lansia")
	}
}
