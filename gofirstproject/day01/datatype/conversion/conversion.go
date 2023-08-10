package main

import (
	"fmt"
	"math"
	"strconv"
)

//conversion == integer to float, float to integer
//parsing == string ke int sama float, boolean ke string

// convert integer to float
func intToFloat() {
	x := 315
	var d float64 = float64(x)
	println(d, x)
	fmt.Printf("Int to Float %f", d)      //% untuk inject
	fmt.Printf("\nInt to Float %8.2f", d) //%8 panjangnya float. 2f untuk angka dibelakang koma

	fmt.Printf("Integer base 10 = %d", x) //%d untuk integer
	fmt.Printf("\nOctal = %b", d)         //%b untuk octal

	//format alternatif - string to float
	format := strconv.FormatFloat(d, 'f', 2, 64) //2 dan 64, digit di belakang koma dan type float64
	println("format : ", format)

	//another method func SprintF, angka ke string
	sprintFormat := fmt.Sprintf("SprintF : %8.2f", d)
	println(sprintFormat)

}

func floatToInt() {
	f := 509.98
	i := int64(f)
	formatInt := strconv.FormatInt(i, 10)
	fmt.Println("\nFloat to Int 1 : ", formatInt)
	fmt.Println("Float to Int 2 : ", i)
}

func intToString() {
	price := 45
	totalPrice := 198.78
	//concat
	output := "Total Price : " + strconv.Itoa(price) //itoa conversion int to string
	//ditampung bisa 2 cara
	// temp := int(totalPrice)
	// total := "Total Price Temp : " + strconv.Itoa(temp)
	total := "Total Price : " + strconv.Itoa(int(totalPrice))

	fmt.Println(output)
	fmt.Println(total)
}

func totalPrice() {
	price := 98.9
	qty := 5
	totalPrice := price * float64(qty)
	fmt.Println("Total Price : ", totalPrice)
	fmt.Println("Total Price Pembulatan : ", math.Round(totalPrice)) //pembulatan Round
}

// contoh parsing ke integer
func parseToInt() {
	val1 := "98"
	//if err not error, err return nil
	int1, err := strconv.ParseInt(val1, 10, 64) //10 base, 64 bit sizenya
	if err == nil {                             //nil = tidak error, jika error sama dengan nil, maka seterusnya
		fmt.Println("Parsed value : ", int1)
	} else {
		fmt.Println("Can not value : ", val1, err)
	}

}

// contoh parsing ke float
func parseToFloat() {
	val1 := "98.98"
	//if err not error, err return nil
	int1, err := strconv.ParseFloat(val1, 64) //10 base, 64 bit sizenya
	if err == nil {                           //nil = tidak error, jika error sama dengan nil, maka seterusnya
		fmt.Println("Parsed value : ", int1)
	} else {
		fmt.Println("Can not value : ", val1, err)
	}

}

func parseStringToBool() {
	val1 := "true"
	val2 := "false"
	val3 := "selain kalimat bawaan, auto false"

	bool1, err1 := strconv.ParseBool(val1)
	bool2, err2 := strconv.ParseBool(val2)
	bool3, err3 := strconv.ParseBool(val3)
	if err1 == nil {
		fmt.Println("Bool1 :", bool1)
	} else {
		fmt.Println("Error Parse Bool1 :", bool1, err1)
	}
	if err2 == nil {
		fmt.Println("Bool2 :", bool2)
	} else {
		fmt.Println("Error Parse Bool2 :", bool2, err2)
	}
	if err3 == nil {
		fmt.Println("Bool3 :", bool3)
	} else {
		fmt.Println("Error Parse Bool3 :", bool3, err3)
	}

}

func main() {
	// intToFloat()
	// floatToInt()
	// intToString()
	// totalPrice()
	parseToInt()
	// parseToFloat()
	// parseStringToBool()
}
