package main

import "fmt"

func basicPattern() {
	for col := 0; col < 5; col++ { //col > column, cetak horizontal
		fmt.Print("* ")
	}
	for row := 0; row < 5; row++ { //row > row, cetak vertical
		fmt.Println("*")
	}
}

func bintangKotak(count int) { //parameter
	//looping cetak baris, ke bawah
	for row := 0; row < count; row++ { //passing lewat parameter

		//cetak looping kolom, ke samping
		for col := 0; col < count; col++ {
			fmt.Print("* ")
		}
		fmt.Println() //pindah baris, untuk mencetak ulang bintang horizontal
	}
}

func bintangKotakNoPara() { //parameter
	//looping cetak baris, ke bawah
	for row := 0; row < 5; row++ { //passing lewat parameter

		//cetak looping kolom, ke samping
		for col := 0; col < 5; col++ {
			fmt.Print("* ")
		}
		fmt.Println() //pindah baris, untuk mencetak ulang bintang horizontal
	}
}

func bintangKotakIndex(count int) { //parameter
	//looping cetak baris, ke bawah
	for row := 0; row < count; row++ { //passing lewat parameter

		//cetak looping kolom, ke samping
		for col := 0; col < count; col++ {
			fmt.Printf("[%d:%d] ", row, col)
		}
		fmt.Println() //pindah baris, untuk mencetak ulang bintang horizontal
	}
}

func bintangLeftDiagonal() {
	//horizontal
	// 00, 11, 22, 33, 44 >> index
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if row == col {
				fmt.Printf("[%d-%d] ", row, col)
			} else {
				fmt.Print(" * ")
			}
		}
		fmt.Println()
	}
}

func bintangKosongTengah() {
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if row == 0 || row == 5-1 || col == 0 || col == 5-1 {
				fmt.Printf("[%d-%d] ", row, col)

			} else {
				fmt.Print("***** ")
			}
		}
		fmt.Println()
	}
}

func bintangSegitigaSiku() {
	for row := 0; row <= 5; row++ {
		for col := 0; col <= row; col++ {
			fmt.Printf("[%d-%d] ", row, col)
		}
		fmt.Println()
	}
}

func bintangSegitigaSikuKanan() {
	for row := 0; row < 5; row++ {
		for col := 0; col < 4-row; col++ {
			fmt.Print("***** ")
		}
		for cols := 0; cols <= row; cols++ {
			fmt.Printf("[%d-%d] ", row, cols)
		}
		fmt.Println()
	}
}

func bintangSegitigaSikuKananDif() {
	for row := 1; row <= 5; row++ {
		for col := 1; col <= 5-row; col++ {
			fmt.Print("***** ")
		}
		for cols := 1; cols <= row; cols++ {
			fmt.Printf("[%d-%d] ", row, cols)
		}
		fmt.Println()
	}
}

func bintangSegitigaSisi() {
	for row := 1; row <= 5; row++ {
		for col := 1; col <= 5-row; col++ {
			fmt.Print("***** ")
		}
		for cols := 1; cols <= row; cols++ {
			fmt.Printf("[%d-%d] ", row, cols)
		}
		fmt.Println()
	}
}

func main() {
	// basicPattern()
	// bintangKotak(5)
	// bintangKotakNoPara()
	// bintangKotakIndex(5)
	// bintangLeftDiagonal()
	// bintangKosongTengah()
	// bintangSegitigaSiku()
	// bintangSegitigaSikuKanan()
	// bintangSegitigaSikuKananDif()
}
