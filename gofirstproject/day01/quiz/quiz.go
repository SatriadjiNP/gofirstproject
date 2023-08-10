package main

import (
	"fmt"
	"math/rand"
)

func quiz1() {
	var x int
	var y int
	var k int
	fmt.Println("Input Number : ")
	fmt.Scan(&x, &y, &k)
	//x = posisi awal, y = target posisi, k = jarak tempuh satu kali lompat
	mod := (y - x) % k
	total := (y - x) / k
	if mod != 0 {
		fmt.Println(total + 1)
	} else {
		fmt.Println(total)
	}

}

func quiz2() {
	// 1 jam = 60 menit, 1 hari = 24 jam, 1 hari = 24 x 60 = 1440 menit, 1 tahun = 365 hari = 525600
	var minute int
	fmt.Println("Input Menit : ")
	fmt.Scan(&minute)
	year := minute / 525600
	mod := minute % 525600
	day := mod / 1440
	mod = mod % 1440

	fmt.Printf("Output : %d year %d day ", year, day)
}

func quiz3() {
	var jarak, bensin, harga, total_harga float64

	fmt.Println("Masukkan jarak yang ditempuh : ")
	fmt.Scan(&jarak)

	fmt.Println("Masukkan bensin per kilometer yang digunakan : ")
	fmt.Scan(&bensin)

	fmt.Println("Masukkan harga bensin per liter : ")
	fmt.Scan(&harga)

	total_harga = jarak / bensin * harga
	fmt.Printf("Uang yang harus dikeluarkan untuk menempuh jarak %8.2f = Rp.%8.2f ", jarak, total_harga)
}

func quiz4() {
	var tahun_lahir int
	fmt.Println("Masukkan Tahun Lahir : ")
	fmt.Scan(&tahun_lahir)

	switch shio := tahun_lahir; {
	case shio%12 == 0:
		fmt.Println("Anda lahir di tahun Monyet")
	case shio%12 == 1:
		fmt.Println("Anda lahir di tahun Ayam")
	case shio%12 == 2:
		fmt.Println("Anda lahir di tahun Anjing")
	case shio%12 == 3:
		fmt.Println("Anda lahir di tahun Babi")
	case shio%12 == 4:
		fmt.Println("Anda lahir di tahun Tikus")
	case shio%12 == 5:
		fmt.Println("Anda lahir di tahun Kerbau")
	case shio%12 == 6:
		fmt.Println("Anda lahir di tahun Harimau")
	case shio%12 == 7:
		fmt.Println("Anda lahir di tahun Kelinci")
	case shio%12 == 8:
		fmt.Println("Anda lahir di tahun Naga")
	case shio%12 == 9:
		fmt.Println("Anda lahir di tahun Ular")
	case shio%12 == 10:
		fmt.Println("Anda lahir di tahun Kuda")
	case shio%12 == 11:
		fmt.Println("Anda lahir di tahun Kambing")
	default:
		fmt.Println("TANYA MAMAK KAU, KAU LAHIR TAHUN APA?")
	}

}

// simulasiAtm
// (Simulasi ATM). BUat program simulasi ATM, dimana user menginput pin password pertama kali.
// jika pin password benar, maka dilanjutkan ke input berikutnya yaitu jumlah uang yang akan diambil.
// ATM hanya memiliki pecahan $10, $5 dan $1. Misal user input uang yang akan diambil $36,
// maka program akan menampilkan output : "$10 = 3 lembar, $5 = 1 lembar, $1 lembar".
// jika pin salah, maka tampilkan "kesempatan 2 kali lagi untuk mencoba". JIka ketiga kalinya salah, tampilkan "pin Anda di lock".
func quiz5() {
	/* fmt.Println("Enter pin : ")
	var pin int
	_, err := fmt.Scan(&pin)
	if err != nil {
		log.Fatal(err)
	}
	pinBenar := 0000
	if pin == pinBenar {
		fmt.Println("Enter jumlah uang : ")
		var jumlah int
		_, err := fmt.Scan(&jumlah)
		if err != nil {
			log.Fatal(err)
		}
		//menghitung uang yang diambil
		pecahan10 := jumlah / 10
		sisaUang := jumlah % 10
		pecahan5 := sisaUang / 5
		sisaUang = sisaUang % 5
		pecahan1 := sisaUang / 1
		fmt.Printf("$10 = %d lembar, $5 = %d lembar, $1 = %d", pecahan10, pecahan5, pecahan1)
	} else {
		fmt.Println("Pin salah (kesempatan 2 kali lagi untuk mencoba) : ")
		var pin int
		_, err := fmt.Scan(&pin)
		if err != nil {
			log.Fatal(err)
		}
		if pin == pinBenar {
			fmt.Println("Enter jumlah uang : ")
			var jumlah int
			_, err := fmt.Scan(&jumlah)
			if err != nil {
				log.Fatal(err)
			}
			//menghitung jumlah uang
			pecahan10 := jumlah / 10
			sisaUang := jumlah % 10
			pecahan5 := sisaUang / 5
			sisaUang = sisaUang % 5
			pecahan1 := sisaUang / 1
			fmt.Printf("$10 = %d lembar, $5 = %d lembar, $1 = %d", pecahan10, pecahan5, pecahan1)
		} else {
			fmt.Println("Pin salah (kesempatan 1 kali lagi untuk mencoba) : ")
			var pin int
			_, err := fmt.Scan(&pin)
			if err != nil {
				log.Fatal(err)
			}
			if pin == pinBenar {
				fmt.Println("Enter jumlah uang : ")
				var jumlah int
				_, err := fmt.Scan(&jumlah)
				if err != nil {
					log.Fatal(err)
				}
				pecahan10 := jumlah / 10
				sisaUang := jumlah % 10
				pecahan5 := sisaUang / 5
				sisaUang = sisaUang % 5
				pecahan1 := sisaUang / 1

				fmt.Printf("$10 = %d lembar, $5 = %d lembar, $1 = %d", pecahan10, pecahan5, pecahan1)
			} else {
				fmt.Println("Pin Anda di lock")
			}
		}
	} */

	var pin, jumlah_uang int

	fmt.Println("Inputkan pin Passsword Anda : ")
	fmt.Scan(&pin)
	if pin != 000000 {
		fmt.Println("Pin Salah ! Kesempatan 2 Kali Lagi Untuk Mencoba : ")
		fmt.Println("Inputkan pin Passsword Anda : ")
		fmt.Scan(&pin)
		if pin != 000000 {
			fmt.Println("Pin Salah ! Kesempatan 1 Kali Lagi Untuk Mencoba : ")
			fmt.Println("Inputkan pin Passsword Anda : ")
			fmt.Scan(&pin)
			if pin != 000000 {
				fmt.Println("Pin Salah pin anda di lock : ")
			} else {
				fmt.Printf("Masukkan Nominal Uang yang akan diambil : ")
				fmt.Scan(&jumlah_uang)
				nominal1 := jumlah_uang / 10
				nominal2 := (jumlah_uang % 10) / 5
				nominal3 := jumlah_uang % 5
				fmt.Printf("$10 = %d Lembar, $5 = %d Lembar, $1 = %d Lembar", nominal1, nominal2, nominal3)
			}
		} else {
			fmt.Printf("Masukkan Nominal Uang yang akan diambil : ")
			fmt.Scan(&jumlah_uang)
			nominal1 := jumlah_uang / 10
			nominal2 := (jumlah_uang % 10) / 5
			nominal3 := jumlah_uang % 5
			fmt.Printf("$10 = %d Lembar, $5 = %d Lembar, $1 = %d Lembar", nominal1, nominal2, nominal3)
		}
	} else {
		fmt.Printf("Masukkan Nominal Uang yang akan diambil : ")
		fmt.Scan(&jumlah_uang)
		nominal1 := jumlah_uang / 10
		nominal2 := (jumlah_uang % 10) / 5
		nominal3 := jumlah_uang % 5

		fmt.Printf("$10 = %d Lembar, $5 = %d Lembar, $1 = %d Lembar", nominal1, nominal2, nominal3)
	}
}

//TebakHadiah. Buat program permainan tebak berhadiah. user input dua digit angka integer dalam skala (10-99).
// computer akan mengacak jawaban dalam dua digit. Contoh:
// User input : 44
// Komputer random : 45
// Display output : exact match : you win Rp. 50.000
// (jika ada satu angka yang cocok, maka tampilkan seperti di atas)

// User input : 44
// Komputer random : 44
// Display output : exact match : you win Rp. 100.000
// (jika ada satu angka yang cocok, maka tampilkan seperti di atas)

// User input : 44
// Komputer random : 85
// Display output : You Lose !
func quiz6() {
	// var angka int
	// fmt.Println("Masukkan angka : ")
	// fmt.Scan(&angka)

	// num1 := angka / 10
	// sisa := angka % 10
	// num2 := sisa % 10

	// random := rand.Intn(100) + 11
	// comp_random := random / 10
	// sisa_random := random % 10
	// comp_random1 := sisa_random % 10

	// number_satu := strconv.Itoa(num1)
	// number_dua := strconv.Itoa(num2)
	// number_tiga := number_satu + number_dua
	// gacha1 := strconv.Itoa(random)
	// gacha2 := strconv.Itoa(comp_random)
	// gacha3 := strconv.Itoa(comp_random1)

	// if gacha1 == number_satu {
	// 	println("User input : ", angka)
	// 	println("Komputer random : ", gacha1)
	// 	println("Match all digit : you win Rp.100.000")
	// } else if gacha2 == number_dua {
	// 	println("User input : ", angka)
	// 	println("Komputer random : ", gacha1)
	// 	println("Exact match : you win Rp.50.000")
	// } else if gacha3 == number_tiga {
	// 	println("User input : ", angka)
	// 	println("Komputer random : ", gacha1)
	// 	println("Exact match : you win Rp.50.000")
	// } else {
	// 	println("User input : ", angka)
	// 	println("Komputer random : ", gacha1)
	// 	println("You Lose!")
	// }

	var gacha, angka1, angka2, angka3, angka4 int
	fmt.Println("$$$ ==> SELAMAT DATANG DI PERMAINAN TEBAK BERHADIAH <=== $$$")
	fmt.Println("")
	fmt.Println("Silahkan Masukkan 2 digit Nomor Keberuntungan anda antara 10 - 99")
	fmt.Scan(&gacha)
	angka1 = gacha / 10
	angka2 = gacha % 10

	min := 10
	max := 99
	komputer_random := rand.Intn(max-min) + min
	angka3 = komputer_random / 10
	angka4 = komputer_random % 10

	if gacha < 10 || gacha > 99 {
		fmt.Printf("Input Harus berada pada range 10 - 99")
	} else if gacha == komputer_random {
		fmt.Println("User Input :", gacha)
		fmt.Println("Komputer random :", komputer_random)
		fmt.Printf("Match All Digit : You Win Rp 100.000")
	} else if angka1 == angka3 || angka1 == angka4 || angka2 == angka3 || angka2 == angka4 {
		fmt.Println("User Input :", gacha)
		fmt.Println("Komputer random :", komputer_random)
		fmt.Printf("Exact match : You win Rp 50.000")
	} else {
		fmt.Println("User Input :", gacha)
		fmt.Println("Komputer random :", komputer_random)
		fmt.Println("You Lose !")
	}
}

func main() {
	quiz1()
	// quiz2()
	// quiz3()
	// quiz4()
	// quiz5()
	// quiz6()
}
