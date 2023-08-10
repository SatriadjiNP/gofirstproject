package main

import (
	"fmt"
	"reflect"
	"sort"
)

// array [5] > ada isi panjang indexnya
// slice [] > bebas ga ditentukan bisa

// init array slice == arraylist <>
// init map slice == maplist <>
func initSlice() {
	// init slice with literal value
	list := []int{1, 2, 3, 4, 5}
	fmt.Println(list)

	//using method to create a slice array
	tech := make([]string, 3, 5) //capacitynya langsung dikali 2, bila mau extend index lebih
	tech[0] = "Go"
	tech[1] = "Java"
	tech[2] = "Python"
	//contoh menambahkan lebih dari length index
	// tech[3] = "C#" //dan slice array auto menambahkan panjang length, with append

	//gunakan method append to add new element
	tech = append(tech, "c#", "PHP", "SQL", "MSSQL", "MySQL") //bisa langsung multiple

	fmt.Println(tech)
}

func appendOneSlice() {
	tech := make([]string, 3, 5) //capacitynya langsung dikali 2, bila mau extend index lebih
	tech[0] = "Go"
	tech[1] = "Java"
	tech[2] = "Python"

	//menambahkan framework array ke dalam tech array
	framework := []string{"Springboot", "Flutter"}
	tech = append(tech, framework...)

	fmt.Println(tech)
}

func createSliceFromArray() {
	techs := [6]string{"Java", "Golang", "C#", "Oracle", "Postgre", "Mongodb"}

	//memotong index, 0 = index kecil, 3 = terbesar
	prog := techs[0:3]    //java golang c#
	database := techs[3:] //oracle postgre mongodb

	//merubah database, maka induk dari array juga ikut berubah, techs "Postgre" > "MSSQL", tipe references
	database[1] = "MSSSQL"

	fmt.Println(prog)
	fmt.Println(database)
	fmt.Println(techs)

}

func copySlice() {
	techs := []string{"Java", "Golang", "C#", "Oracle", "Postgre", "Mongodb"}
	//memotong db, dan menampung techs, oracle postgre, mongodb (fetch)
	pl := techs[3:]
	//membuat slice variable db
	db := make([]string, 3)
	// use method copy, args: destination, arg2: source
	copy(db, pl)
	db[1] = "SQL"

	fmt.Println(db)
}

func replaceElementSlice() {
	techs := []string{"Java", "Golang", "C#", "Oracle", "Postgre", "Mongodb"}

	replaceElement := []string{"Vb", "Pascal"}

	// menganti java golang = vb pascal
	copy(techs[0:2], replaceElement)
	fmt.Println(techs)
}

func sortSlice() {
	techs := []string{"Java", "Golang", "C#", "Oracle", "Postgre", "Mongodb"}

	// sorting
	sort.Strings(techs)

	fmt.Println(techs)

	for _, v := range techs[0:3] {
		fmt.Print(v, " ")

	}
}

func displayArray(arr [3]string) {

}

func compareArray() {
	emp1 := [3]string{"A", "B", "C"}
	emp2 := [3]string{"A", "B", "C"}

	// output false or true, untuk menemukan sama atau tidaknya element
	compare := emp1 == emp2

	fmt.Println("Result compare :", compare)
}

func compareSlice() {
	emp1 := []string{"A", "B", "C"}
	emp2 := []string{"A", "E", "D"}

	// output false or true, untuk menemukan sama atau tidaknya element
	// cara menggunakan slice, gunakan method reflect
	compare := reflect.DeepEqual(emp1, emp2)

	fmt.Println("Result compare :", compare)
}

func main() {
	// initSlice()
	// appendOneSlice()
	// createSliceFromArray()
	// copySlice()
	// replaceElementSlice()
	// sortSlice()
	// displayArray(emp)
	// compareArray()
	compareSlice()
}
