package main

import "fmt"

func aVeryBigSum(ar []int64) int64 {
	// Write your code here

	var sum int64

	for _, num := range ar {
		sum += num
	}
	return sum
}

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var arr1, arr2 int32

	size := len(arr)

	for i := 0; i < size; i++ {
		arr1 += arr[i][i]
		arr2 += arr[i][size-i-1]
	}

	diff := (arr1 - arr2)
	if diff < 0 {
		return -diff
	}

	return diff
}

func main() {
	fmt.Println("Soal aVeryBigSum")
	fmt.Println(aVeryBigSum([]int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}))

	println()

	fmt.Println("Soal diagonalDifferent")
	matrix := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	result := diagonalDifference(matrix)
	fmt.Println(result) // Output: 2
}
