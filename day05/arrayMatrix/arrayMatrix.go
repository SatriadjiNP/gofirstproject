package main

import (
	"fmt"
)

func initRow() {
	matrix := [1][5]int{
		{1, 2, 3, 4, 5},
	}
	fmt.Println(matrix)
}

func initCol() {
	matrix := [5][1]int{
		{1},
		{2},
		{3},
		{4},
		{5},
	}
	fmt.Println(matrix)
}

// input dengan 2 function
func lowerTriangle() [][]int {
	matrix := [][]int{
		{1, 0, 0},
		{2, 1, 0},
		{3, 1, 1},
	}
	return matrix
}

// untuk memanggil matrix
func printMatrix(matrix [][]int) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			// fmt.Print(matrix[row][col], " ")
			fmt.Printf("%d ", matrix[row][col])
		}
		fmt.Println()
	}
}

// contoh print 1 function
func initArrayValue() [5][5]int {
	// default zero value, need var
	var matrix [5][5]int
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			matrix[row][col] = row + col
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

func initDiagonal() [5][5]int {
	var matrix [5][5]int
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if row == col {
				matrix[row][col] = 10
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
	return matrix
}

// contoh matrix dari google
func initMatrix() {
	// Membuat matriks 2x3
	matrix := make([][]int, 3)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, 3)
	}

	// Mengisi matriks dengan nilai
	// row col [][]
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6
	matrix[2][0] = 7
	matrix[2][1] = 8
	matrix[2][2] = 9

	// Menampilkan matriks
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}

func main() {
	// initMatrix()
	// initRow()
	// initCol()

	// func harus udah ada value indexnya, di set manual
	// printMatrix(lowerTriangle())

	// initArrayValue()
	initDiagonal()
}
