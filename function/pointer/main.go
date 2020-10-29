package main

import "fmt"

// pointer di function
// digunakan ketika ingin merubah data yang ada didalam function lain
type Address struct {
	City, Province, Country string
}

// Before
// func changeCountryToIndonesia(address Address) {
// 	address.Country = "Indonesia" // data tidak akan berubah
// }

// func main() {
// 	var alamat = Address{
// 		City:     "Subang",
// 		Province: "Jawa Barat",
// 		Country:  "_",
// 	}
// 	changeCountryToIndonesia(alamat)
// 	fmt.Println(alamat) // data tidak akan berubah
// }

// After
func changeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var alamat = Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "_",
	}
	changeCountryToIndonesia(&alamat)
	fmt.Println(alamat) // data akan berubah
}
