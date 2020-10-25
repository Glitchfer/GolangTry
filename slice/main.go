package main

import "fmt"

func main() {
	// kurung buka tutup dengan "..." dan angka merupakan array
	var month = [...]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
	// sedangkan kurung dengan isi kosong "( []string{"bla","bblab"} )"adalah slice

	// Memilih potongan array dari index ke 4 hingga sebelum 7
	var slice1 = month[4:7]
	fmt.Println(slice1)

	// merubah array
	month[5] = "Diubah"
	fmt.Println(month)
	fmt.Println(slice1)

	// merubah isi slice = merubah data array
	slice1[0] = "Loncat"
	fmt.Println(month)
	fmt.Println(slice1)

	// Melihat kapasitas array
	fmt.Println(cap(slice1))

	fmt.Println("panjang array slice1 adalah", len(slice1))
	fmt.Println("panjang array month adalah", len(month))

	var slice2 = month[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Eko")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)
}
