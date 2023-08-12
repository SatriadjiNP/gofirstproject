package main

import (
	"fmt"
	"log"
	"math/rand"
)

func oddOrEven(number int) {
	if number%2 == 0 {
		println("Even Number")
	} else {
		println("Odd Number")
	}
}

/*
	fizzbuzz

number = bisa dibagi 3 & 5 -> fizzbuzz
number = bisa dibagi 3 -> fizz
number = bisa dibagi 5 -> buzz
number = ga bisa dibagi 3 & 5 -> unknown/tidak diketahui
*/
func fizzBuzz(number int) {
	if number%3 == 0 && number%5 == 0 {
		println("Fizzbuzz")
	} else if number%3 == 0 {
		println("Fizz")
	} else if number%5 == 0 {
		println("Buzz")
	} else {
		println("Unknown/tidak diketahui")
	}
}

func randomNumber() {
	// range ouput 0 - 9
	random := rand.Intn(10)
	println("Random : ", random)

	// range ouput : 1 - 10
	/* min := 1
	max := 1000
	randomRange := rand.Intn(max-min) + min */
	randomRange := rand.Intn(10) + 10
	println("Random Range : ", randomRange)
}

func initialIf() {
	random := rand.Intn(10)
	if random < 5 {
		println("low")
	} else if random >= 5 {
		println("high")
	}

	//gaya GO penulisan if
	// if random := rand.Intn(10); random < 5 {
	// 	println("low")
	// } else if random >= 5 {
	// 	println("high")
	// }
}

func guestNumber() {
	angka1 := rand.Intn(10)
	angka2 := rand.Intn(6)
	angka3 := angka1 - angka2

	fmt.Printf("Berapa hasil %d - %d = ?\n", angka1, angka2)
	// println("Berapa hasil dari 7 - 3 ? = ")
	//declare var jawaban untuk nampung inputan dari console
	var jawaban int
	_, err := fmt.Scan(&jawaban) //& menampung jawaban dari console //_ (underscore) adalah jumlah inputan
	if err != nil {
		log.Fatal(err)
	}
	if jawaban == angka3 {
		println("Jawaban Anda Benar")
	} else {
		println("Jawaban Anda Salah")
	}
	// println("Jawaban : ", jawaban)
}

func multipleArgs() {
	fmt.Println("Input 3 angka : ")
	var angka1, angka2, angka3 int
	_, err := fmt.Scan(&angka1, &angka2, &angka3) //bila mao tampilin
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Jawaban : ", angka1, angka2, angka3)
}

// switch expression
// case condition
func tebakHari(day int) {
	switch n := day; {
	case n == 0:
		fmt.Println("Sunday")
	case n == 1:
		fmt.Println("Monday")
	case n == 2:
		fmt.Println("Tuesday")
	case n == 3:
		fmt.Println("Wednesday")
	case n == 4:
		fmt.Println("Thursday")
	case n == 5:
		fmt.Println("Friday")
	case n == 6:
		fmt.Println("Saturday")
	default:
		fmt.Println("Not Today")
	}
}

func main() {
	// oddOrEven(8)
	// fizzBuzz(30)
	// randomNumber()
	// initialIf()
	// guestNumber()
	// multipleArgs()
	// tebakHari(8)
}
