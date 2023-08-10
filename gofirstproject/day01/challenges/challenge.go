package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

// luas lingkaran
func soal1(number float64) {
	phi := 3.1416
	luas := phi * number * number
	fmt.Println("luas lingkaran :", luas)

}

// sudut segitiga
func soal2(numb1, numb2 int) {
	jumlah := 180
	numb3 := jumlah - numb1 - numb2

	println("sudut1 :", numb1)
	println("sudut2 :", numb2)
	println("sudut3 :", numb3)

}

// rata-rata
func soal3(angka1, angka2, angka3 float64) {
	bagi := 3
	average := (angka1 + angka2 + angka3) / float64(bagi)
	fmt.Println(average)

}

// tahun kabisat
func soal4(kabisat int) {
	if (kabisat%4 == 0 && kabisat%100 != 0) || (kabisat%400 == 0) {
		println(kabisat, "Tahun Kabisat")
	} else {
		println(kabisat, "Bukan Tahun Kabisat")
	}
}

// detik ke menit
func soal5(detik int) {
	menit := detik / 60
	detikSisa := detik % 60

	fmt.Printf("Result : %d menit %d detik\n", menit, detikSisa)
}

// detik ke hari
func soal6(second int) {
	day := second / 86400
	dayleft := second % 86400
	hour := dayleft / 3600
	hourleft := dayleft % 3600
	minute := hourleft / 60
	seconds := hourleft % 60

	fmt.Printf("Hasil Konversi : %d day %d hour %d minute %d second ", day, hour, minute, seconds)

	// var hari, jam, menit, detik, sisa int
	// hari = second / 86400
	// sisa = second % 86400
	// jam = sisa / 3600
	// sisa = sisa % 3600
	// menit = sisa / 60
	// detik = sisa % 60

	// fmt.Printf("Hasil Konversi : %d hari %d jam %d menit %d detik ", hari, jam, menit, detik)
}

// reverse number
func soal7(number int) {
	var angka1, angka2, angka3, angka4, sisa int
	if number < 1000 || number > 10000 {
		fmt.Println("Input number range >= 1000 dan < 10000")
	} else {
		angka1 = number / 1000
		sisa = number % 1000
		angka2 = sisa / 100
		sisa = sisa % 100
		angka3 = sisa / 10
		angka4 = sisa % 10

		reverse := strconv.Itoa(angka4) + strconv.Itoa(angka3) + strconv.Itoa(angka2) + strconv.Itoa(angka1)
		fmt.Printf("Reverse : %d -> %s ", number, reverse)
	}

	// if number >= 1_000 && number < 10_000 {
	// 	digit1 := number % 10
	// 	digit2 := (number / 10) % 10
	// 	digit3 := (number / 100) % 10
	// 	digit4 := number / 1000

	// 	reversed := (digit1 * 1000) + (digit2 * 100) + (digit3 * 10) + digit4

	// 	fmt.Printf("\n(7.) Reverse Number : %d --> %d", number, reversed)
	// } else {
	// 	fmt.Printf("\n(7.) Input Number harus dalam Range 1_000 - 10_000")
	// }
}

// Sum Number
func soal8(number int) {
	var angka1, angka2, angka3, angka4, sisa int
	if number < 1000 || number > 10000 {
		fmt.Println("Input number range >= 1000 dan < 10000")
	} else {
		angka1 = number / 1000
		sisa = number % 1000
		angka2 = sisa / 100
		sisa = sisa % 100
		angka3 = sisa / 10
		angka4 = sisa % 10

		jumlah := angka1 + angka2 + angka3 + angka4
		fmt.Printf("Sum Number : %d -> %d ", number, jumlah)
	}
}

// sort tiga angka
func soal9() {
	var num1, num2, num3 int
	fmt.Printf("Masukkan 3 angka : \n")
	_, err := fmt.Scan(&num1, &num2, &num3)

	if num1 < num2 && num1 < num3 && num2 < num3 {
		fmt.Printf("%d , %d, %d", num1, num2, num3)
	} else if num1 > num2 && num1 < num3 && num2 < num3 {
		fmt.Printf("%d, %d, %d", num2, num1, num3)
	} else if num1 > num2 && num1 > num3 && num2 < num3 {
		fmt.Printf("%d, %d, %d", num2, num3, num1)
	} else if num1 < num2 && num1 > num3 && num2 > num3 {
		fmt.Printf("%d, %d, %d", num3, num1, num2)
	} else if num1 > num2 && num1 > num3 && num2 > num3 {
		fmt.Printf("%d, %d, %d", num3, num2, num1)
	} else if num1 < num2 && num1 < num3 && num2 > num3 {
		fmt.Printf("%d, %d, %d", num1, num3, num2)
	} else if err != nil {
		log.Fatal(err)
	}
}

// guntingbatukertas
func soal10() {
	//inputan untuk pilih suit
	var suit int
	fmt.Printf("Pilih : (1)Gunting  (2)Batu  (3)Kertas\n")
	fmt.Printf("Masukkan suit Anda : ")
	user, err := fmt.Scan(&suit)

	pil1 := "Gunting"
	pil2 := "Batu"
	pil3 := "Kertas"
	min := 1
	max := 3

	//random computer, %s string
	computer := rand.Intn(max-min) + min
	if user == 1 && computer == 1 {
		fmt.Printf("Anda Seri, computer [%s] vs Anda [%s]\n", pil1, pil1)
	} else if user == 1 && computer == 2 {
		fmt.Printf("Anda Kalah, computer [%s] vs Anda [%s]\n", pil1, pil2)
	} else if user == 1 && computer == 3 {
		fmt.Printf("Anda Menang, computer [%s] vs Anda [%s]\n", pil1, pil3)
	} else if user == 2 && computer == 1 {
		fmt.Printf("Anda Menang, computer [%s] vs Anda [%s]\n", pil2, pil1)
	} else if user == 2 && computer == 2 {
		fmt.Printf("Anda Seri, computer [%s] vs Anda [%s]\n", pil2, pil2)
	} else if user == 2 && computer == 3 {
		fmt.Printf("Anda Kalah, computer [%s] vs Anda [%s]\n", pil2, pil3)
	} else if user == 3 && computer == 1 {
		fmt.Printf("Anda Kalah, computer [%s] vs Anda [%s]\n", pil3, pil1)
	} else if user == 3 && computer == 2 {
		fmt.Printf("Anda Menang, computer [%s] vs Anda [%s]\n", pil3, pil2)
	} else if user == 3 && computer == 3 {
		fmt.Printf("Anda Seri, computer [%s] vs Anda [%s]\n", pil3, pil3)
	} else if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// soal1(10)
	// soal2(30, 60)
	// soal3(15, 15, 25)
	// soal4(2002)
	// soal5(12345)
	// soal6(1000000)
	// soal7(1234)
	// soal8(1234)
	// soal9()
	soal10()
}
