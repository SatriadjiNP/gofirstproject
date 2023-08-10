package main

import (
	"fmt"
	"strings"
	"unicode"
)

// memotong dari index 1 menjadi (ello) karena [1:] > motong dan mulai dari index 1, huruf E
// [1:5] 1 > index ke 1, dan 4 sama dengan panjang index yang ditampilkan
// (tapi panjang lenght mulai dari 1 bukan 0).. ELLO
func wordString() {
	word := "Hello World" //index dimulai dari 0 > huruf H

	fmt.Println("Length : ", len(word))                   //menghitung panjang huruf
	fmt.Println("To Uppercase : ", strings.ToUpper(word)) //ubah jadi KAPITAL

	ello := word[1:5]
	world := word[6:11]
	foundW := strings.Index(word, "W")
	fmt.Println("Hello Substr : ", ello)
	fmt.Println("Found W at position : ", foundW)
	fmt.Println("World : ", world)

}

func basicLooping() {
	counter := 0
	for {
		fmt.Println("Counter : ", counter) //mencetak counter
		counter++                          //tambah 1 dan looping sampai 5
		if counter == 5 {
			break //berenti, dan keluar program
		}
	}
}

func conditionLoop() {
	counter := 0
	for counter <= 5 {
		fmt.Println("Counter : ", counter) //mencetak counter
		counter++                          //tambah 1 dan looping sampai 5
	}
}

func intiLoop() {
	for counter := 0; counter < 5; counter++ {
		if counter == 2 {
			continue //skip index ke 2
			// break
		}
		fmt.Println("Counter : ", counter)

	}
}

func rangeIndexLooping() {
	word := "Golang"
	for index, value := range word { //_, unuseful variable, v ganti index, range wordnya,
		// kalo mao munculin valuenya kasih variable lagi di samping index
		fmt.Println("Index : ", index, value, string(value)) //value ini bertipe data rune/ascii, kalo mao diubah lagi ke string

	}
	for _, value := range word { //tidak mao menampilkan indexnya, jadi diganti pake underscore ( _ )
		fmt.Println("Index : ", value, string(value))

	}
}

// inputan : BootCamp gOlang
func countUpperChar(words string) {
	countUpper := 0
	for _, value := range words {
		if unicode.IsUpper(value) { //bisa pake IsSpace(spasi), IsLower, dan lain-lain
			countUpper++
		}
	}
	fmt.Println(words, "Char Uppercase : ", countUpper)
}

// bilangan prima
func switchInitLooping() {
	for counter := 0; counter < 20; counter++ {
		switch counter / 2 { //switchKotak > expression = counter
		case 2, 3, 5, 7:
			fmt.Println("Prime Number : ", counter/2)
		default:
			fmt.Println("Non Prime Number : ", counter/2)
			break
		}
	}
}

func switchValLooping() {
	for counter := 0; counter < 20; counter++ {
		switch val := counter / 2; val { //switchKotak > expression = counter
		case 2, 3, 5, 7:
			fmt.Println("Prime Number : ", val)
		default:
			fmt.Println("Non Prime Number : ", val)
			break
		}
	}
}

func main() {
	// wordString()
	// basicLooping()
	// conditionLoop()
	// intiLoop()
	// rangeIndexLooping()
	// countUpperChar("BootCamp gOlang")
	// switchInitLooping()
	switchValLooping()
}
