package main

import (
	"fmt"
)

// function initArray with 1 parameter & has return array
func initArray(n int) []int { //parameter with return
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

// soal no 1
func sumAllElementArray(arr []int) int { //with return
	sum := 0
	// var sum int
	for i := 0; i < len(arr); i++ {
		// sum = sum + arr[i]
		sum += arr[i]
	}
	return sum
}

// soal no 2
func maxElement(arr []int) int {
	max := 0
	for _, ele := range arr {
		if max < ele {
			max = ele
		}
	}
	return max
}

// soal no 3
func minElement(arr []int) int {
	max := 0
	min := 0
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
			min = i
		}
	}
	return min
}

// soal no 4
func shifting(arr []int, rotate int) []int {
	// length := len(arr)
	// if length < 2 || rotate%length == 0 {
	// 	return arr
	// }
	// // bila rotate sama modulus length, maka akan
	// rotate = rotate % length
	// return append(arr[rotate:], arr[:rotate]...)

	length := len(arr)
	for i := 0; i < rotate; i++ {
		temp := arr[0]
		copy(arr, arr[1:])
		arr[length-1] = temp
	}
	return arr
}

// soal no 5
func linearSearch(arr []int, rotate int, search int) int {
	length := len(arr)
	indexSearch := 0
	for i := 0; i < rotate; i++ {
		temp := arr[0]
		copy(arr, arr[1:])
		arr[length-1] = temp
	}
	for j := 0; j < len(arr); j++ {
		if arr[j] == search {
			indexSearch = j
		}
	}
	return indexSearch
}

// untuk funsi matrix, memanggil func ini kesemua func, tipe harus sama
func displayArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(i, ", ")
	}

}

func main() {
	// karena menggunakan return, maka...
	// return harus ditampung dulu, variable list untuk penampung
	list := initArray(5)
	displayArray(list)
	fmt.Println()
	displayArray(initArray(10))
	// fmt.Println(list)

	// penampung soal no 1
	// number := []int{1, 2, 3, 4, 5}
	fmt.Println(sumAllElementArray([]int{1, 2, 3, 4, 5}))

	// penampung soal no 2
	// number := []int{89, 76, 10, 99, 10}
	fmt.Println(maxElement([]int{89, 76, 10, 99, 10}))

	// soal no 3
	fmt.Println(minElement([]int{1, 1, 12, 10, 5, 12}))

	// soal no 4
	fmt.Println(shifting([]int{2, 3, 5, 7, 11}, 1))

	// soal no 5
	fmt.Println(linearSearch([]int{2, 3, 5, 7, 11}, 1, 7))
}
