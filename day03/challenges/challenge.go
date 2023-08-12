package main

import "fmt"

// soal no 1
func bintangTengah() {
	for row := 0; row < 5; row++ {
		for col := 5 - 1; col >= 0; col-- {
			if row == col {
				fmt.Print("* ")
			} else if row <= col {
				fmt.Print(" ")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

//soal no 2
func bintangSegitigaKanan() {
	for row := 1; row <= 5; row++ {
		for col := 1; col <= 5-row; col++ {
			fmt.Print("-")
		}
		for cols := 1; cols <= row; cols++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

//-- -- -- -- 10
//-- -- -- 10 20
//-- -- 10 20 20
//-- 10 20 20 20
//10 20 20 20 20
//soal no 3
func bintang1020Kanan() {
	for row := 0; row < 5; row++ {
		for col := 5 - 1; col >= 0; col-- {
			if row == col {
				fmt.Printf("10")
			} else if row <= col {
				fmt.Printf("- ")
			} else {
				fmt.Printf("20")
			}
		}
		fmt.Println()
	}
}

//soal no 4
func bintangDiagonal() {
	for row := 0; row <= 5; row++ {
		for col := 5 - 1; col >= 0; col-- {
			if row == col {
				fmt.Print("10")
			} else if row <= col {
				fmt.Print("21")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

//soal no 5
func bintangFull() {
	for row := 0; row < 5; row++ {
		for col := 5 - 1; col >= 0; col-- {
			if row == col {
				fmt.Print("10")
			} else if row <= col {
				fmt.Print("21")
			} else {
				fmt.Print("20")
			}
		}
		fmt.Println()
	}
}

// soal no 6
// 54321
// 12345
// 54321
// 12345
func inputAngkaReverse() {
	var angka int
	fmt.Print("Masukkan Inputan : ")
	fmt.Println()
	fmt.Scanf("%d", &angka)

	for row := 1; row <= angka; row++ {
		if row%2 == 0 {
			for col := 1; col <= angka; col++ {
				fmt.Print(col)
			}
		} else {
			for col := angka; col >= 1; col-- {
				fmt.Print(col)
			}
		}
		fmt.Println()
	}
}

//soal no 7
// 15151
// 24242
// 33333
// 42424
// 51515
func inputAngkaReverseRow() {
	var angka int
	fmt.Print("Masukkan Inputan : ")
	fmt.Println()
	fmt.Scanf("%d", &angka)

	// for row := 1; row <= angka; row++ {
	// 	for ulang := 0; ulang < angka/2; ulang++ { //membagi 2 inputan

	// 		for col := 1; col <= angka; col++ {
	// 			if col == angka {
	// 				fmt.Print(row)
	// 			} else {
	// 				fmt.Print("")
	// 			}
	// 		}
	// 		for col := angka; col >= 1; col-- {
	// 			if col == 1 {
	// 				fmt.Print(angka - row + 1)
	// 			} else {
	// 				fmt.Print("")
	// 			}
	// 		}
	// 	}
	// 	if angka%2 == 1 { //perulangan inputan bisa ganjil, dari habil pembagian ulang angka diatas
	// 		for col := 1; col <= angka; col++ {
	// 			if col == angka {
	// 				fmt.Print(row)
	// 			} else {
	// 				fmt.Print("")
	// 			}
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	for row := 0; row < angka; row++ {
		for col := 0; col < angka; col++ {
			if col%2 == 0 {
				fmt.Print(row + 1)
			} else {
				fmt.Print(angka - row)
			}
		}
		fmt.Println()
	}
}

//soal no 8
// -2-4-
// 1-3-5
// -2-4-
// 1-3-5
// -2-4-
func inputAngkaMod() {
	var angka int
	fmt.Print("Masukkan Inputan : ")
	fmt.Scanf("%d", &angka)

	for row := 0; row < angka; row++ {
		if row%2 == 0 {
			for col := 1; col <= angka; col++ {
				if col%2 == 0 { //bila mod terpenuhi cetak
					fmt.Print(col)
				} else {
					fmt.Print("-")
				}
			}
		} else {
			for col := 1; col <= angka; col++ {
				if col%2 == 0 { //bila mod tidak terpenuhi cetak (-)
					fmt.Print("-")
				} else {
					fmt.Print(col)
				}
			}
		}
		fmt.Println()
	}
}

func main() {
	bintangTengah()
	bintangSegitigaKanan()
	bintang1020Kanan()
	bintangDiagonal()
	bintangFull()
	inputAngkaReverse()
	inputAngkaReverseRow()
	inputAngkaMod()
}
