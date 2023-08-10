package main

import "fmt"

//global variable
var salary = 1_000 //_ buat memisahkan angka

func shortVariable() {
	//declare variable using := operator
	var a = 10 //untype variable
	x := 10
	z, y := 12, "golang"

	//re-assign value for variable x
	x = 100
	println(a, x, z, y)
}

func varDeclare() {
	//init variable default zero value, var itu bisa diganti value, concs engga bisa diganti value
	var a int       // default value : 0
	var word string // default value : ""

	//variable type
	var x int = 10

	//multiple variable with the same type
	var y, z int = 15, 20

	//multiple variable untype
	var k, str = 45, "golang"

	//declare multiple variable at once
	var (
		q    int
		xy       = 20
		zy   int = 12
		msg      = "golang"
		f, g string
	)

	fmt.Println("a : ", a)
	fmt.Println("x : ", x)
	fmt.Println("y : ", y)
	fmt.Println("z : ", z)
	fmt.Println("str : ", str)
	println(f, g, q, zy, k, xy, word, msg)

}

func main() {
	fmt.Println("Hello Satriadji")
	varDeclare()
	shortVariable()
}
