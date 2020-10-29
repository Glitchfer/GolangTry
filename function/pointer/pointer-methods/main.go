package main

// pointer di struct
// tujuannya untuk merubah data struct yang telah di return oleh function lain
// bila tidak menggunakan pointer, data struct tidak akan berubah
type person struct {
	Name string
}

// Before
// func (man person) Married() {
// 	man.Name = "Mr." + man.Name
// }

// func main() {

// 	data := person{"Rahman"}
// 	data.Married()
// 	fmt.Println(data.Name) //Hasilnya data tidak akan berubah

// }

// After
func (man *person) Married() {
	man.Name = "Mr." + man.Name
}

func main() {

	data := person{"Rahman"}
	data.Married()
	// fmt.Println(data.Name) data akan berubah

}
