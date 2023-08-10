package main

import "fmt"

// variable, panjang index, tipe data, {isi index}
func initArrayString() {
	var students [3]string
	students[0] = "Nol"
	students[1] = "Satu"
	students[2] = "Dua"

	fmt.Println(students)
	// tampilkan index 1
	fmt.Println(students[1])

	yugi := [3]string{"Yugi", "Kaiba", "Joy"}
	// memanggil tapi ganti isi index
	yugi[1] = "Mei"
	fmt.Println(yugi)

	//[...] operator, tergantung dari isi index
	number := [...]string{"Nol", "Satu", "Dua"}
	fmt.Println(number)
	//memanggil func displayArray ke dalam func initArray
	displayArray(number)

}

//function array dengan parameter tipe array juga, parameter harus ada panjang length index
// kalo engga ada length index, tipe data slice digunakan
func displayArray(arr [3]string) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[1], " ")
	}
}

func main() {
	initArrayString()
}
