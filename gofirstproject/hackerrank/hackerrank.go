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

func main() {
	fmt.Println("Soal aVeryBigSum")
	fmt.Println(aVeryBigSum([]int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}))
}
