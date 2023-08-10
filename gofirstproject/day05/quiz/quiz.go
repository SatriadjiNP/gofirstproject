package main

import (
	"fmt"
)

func soal1Diagonal() [5][5]int {
	var matrix [5][5]int
	sum := 0
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if row == col {
				matrix[row][col] = col + 1
				sum += col + 1
			} else if row <= col {
				matrix[row][col] = 20
			} else {
				matrix[row][col] = 21
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println("Total Sum : ", sum)
	return matrix
}

func soal2Diagonal() [5][5]int {
	var matrix [5][5]int
	sum := 0
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if row == col {
				matrix[row][col] = 2 << col
				sum += 2 << col
			} else if row <= col {
				matrix[row][col] = 88
			} else {
				matrix[row][col] = 22
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println("Total Sum : ", sum)
	return matrix
}

// soal no 3
func soal3Siswa() {
	matriks := [9][11]string{
		{"student-0", "A", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
		{"student-1", "D", "B", "A", "B", "C", "A", "E", "E", "A", "D"},
		{"student-2", "E", "D", "D", "A", "C", "B", "E", "E", "A", "D"},
		{"student-3", "C", "B", "A", "E", "D", "C", "E", "E", "A", "D"},
		{"student-4", "A", "B", "D", "C", "C", "D", "E", "E", "A", "D"},
		{"student-5", "B", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
		{"student-6", "B", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
		{"student-7", "E", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
		{"Kunci Jawaban", "D", "B", "D", "C", "C", "D", "A", "E", "A", "D"},
	}

	for row := 0; row < len(matriks)-1; row++ {
		point := 0
		for col := 1; col < len(matriks[row]); col++ {
			if matriks[row][col] == matriks[len(matriks)-1][col] {
				point += 10
			}
		}
		fmt.Printf("\n%s: %d point ", matriks[row][0], point)
	}
}

// soal no 4
func soal4KotakArray() [7][7]int {
	var matrix [7][7]int
	index0 := 17
	midValue := 3
	index1 := index0 + midValue
	index2 := index0 + (midValue * (len(matrix) + 1))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if i+j == len(matrix)-1 && i != j {
				matrix[i][j] = index0
				index0 -= midValue
			} else if i == len(matrix)/2 && j == len(matrix)/2 {
				matrix[i][j] = midValue
			} else if i == 1 && j == len(matrix)-1 {
				matrix[i][j] = midValue * 3
			} else if i > 1 && j == len(matrix)-1 && i != len(matrix)-2 {
				matrix[i][j] = index1
				index1 += 3
			} else if i == len(matrix)-2 && j == len(matrix)-1 {
				matrix[i][j] = midValue * 9
			} else if j > 0 && i == len(matrix)-1 && j != len(matrix)/2 {
				matrix[i][j] = index2
				index2 -= midValue
			} else if i == len(matrix)-1 && j == len(matrix)-4 {
				matrix[i][j] = midValue * 27
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
	return matrix
}

// soal no 5
func soal5IndexSum3() [7][7]int {
	var matriks [7][7]int
	n := 2

	for i := 0; i < len(matriks); i++ {
		for j := 0; j < len(matriks); j++ {
			if i == 0 && j != len(matriks)/2 {
				matriks[i][j] = n
				n += 3
			} else if (i == 0 && j == len(matriks)/2) || (i == len(matriks)-1 && j == len(matriks)/2) || (j == 0 && (i == 1 || i == len(matriks)-2)) || (j == len(matriks)-1 && (i == 1 || i == len(matriks)-2)) {
				matriks[i][j] = 3
			} else if i > 1 && j == len(matriks)-1 && i != len(matriks)-2 {
				matriks[i][j] = n
				n += 3
			} else if i == 0 && j == len(matriks)-1 && i != len(matriks)-2 {
				matriks[i][j] = n
				n += 3
			}
		}
	}
	for row := len(matriks) - 1; row >= 0; row-- {
		for col := len(matriks) - 1; col >= 0; col-- {
			if (row == len(matriks)-1 && col < len(matriks)-1) && col != len(matriks)/2 {
				matriks[row][col] = n
				n += 3
			} else if (row > 1 && col == 0) && row != len(matriks)-2 {
				matriks[row][col] = n
				n += 3
			}
		}

	}
	for i := 0; i < len(matriks); i++ {
		for j := 0; j < len(matriks); j++ {
			fmt.Printf("%d ", matriks[i][j])
		}
		fmt.Println()
	}
	return matriks
}

// soal no 6
func soal6IndexLuar(length int) {
	// var matrix [7][7]int
	matrix := make([][]int, length+1)
	for i := 0; i <= length; i++ {
		matrix[i] = make([]int, length+1)
	}

	// Mengisi nilai pada indeks paling luar
	for row := 0; row < length; row++ {
		for col := 0; col < length; col++ {
			matrix[0][col] = col
			matrix[row][0] = row
			matrix[length-1][col] = col + length - 1
			matrix[row][length-1] = row + length - 1
		}
	}

	// Menampilkan matriks
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			// if matrix[i][j] != 0 {
			fmt.Printf("%d ", matrix[i][j])
			// } else {
			// 	fmt.Print(". ")
			// }
		}
		fmt.Println()
	}
}

// soal no 7
func soal7DiagonalSum(length int) {
	// rows := 7
	// cols := 7
	matrix := make([][]int, length+1)
	for i := 0; i < length+1; i++ {
		matrix[i] = make([]int, length+1)
	}
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			matrix[i][j] = i + j
		}
	}
	var sumBaris int
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			fmt.Print(matrix[i][j], " ")
			sumBaris += matrix[i][j]

		}
		fmt.Print("[", sumBaris, "]")

		sumBaris = 0

		fmt.Println()
	}

	sumKolom := 0
	sumDiagonal := 0
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			sumKolom += matrix[i][j]
			if i == j {
				sumDiagonal += matrix[i][j]
			}

		}
		fmt.Print("[", sumKolom, "]")
		sumKolom = 0
	}
	fmt.Print("[", sumDiagonal, "]")
	// for _, row := range matrix {
	// 	for _, val := range row {
	// 		fmt.Printf("%d ", val)
	// 	}
	// 	fmt.Println()
	// }
}

func main() {
	// soal no 1
	// soal1Diagonal()

	// soal no 2
	// soal2Diagonal()

	// soal no 3
	// soal3Siswa()

	// soal no 4
	// soal4KotakArray()

	// soal no 5
	soal5IndexSum3()

	// soal no 6
	// soal6IndexLuar(7)

	// soal no 7
	// soal7DiagonalSum(7)
}
