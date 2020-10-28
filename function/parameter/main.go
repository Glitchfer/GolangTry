package main

import "fmt"

func main() {

	mahasiswa := make(map[string]string)
	mahasiswa["nama"] = "Ryan"
	mahasiswa["absen"] = "11"

	for key, value := range mahasiswa {

		funcWithParameter(key, value)

	}
}

func funcWithParameter(key string, value string) {
	fmt.Println(key, " anda adalah ", value)
}
