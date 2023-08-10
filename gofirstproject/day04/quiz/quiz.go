package main

import (
	"fmt"
	"strconv"
)

// soal no 1
func evenOrOddNumber(n int) []int {
	arr := make([]int, n)
	if n%2 == 1 {
		arr[0] = 1
	} else if n%2 == 0 {
		arr[0] = 2
	}
	for i := 1; i < len(arr); i++ {
		arr[i] = arr[i-1] + 2
	}
	return arr
}

// soal no 2
func powNumber(n int) []int {
	arr := make([]int, n)
	arr[0] = 1
	for i := 1; i < len(arr); i++ {
		arr[i] = arr[i-1] * 2
	}
	return arr
}

// soal no 3
func fibonnaci(n int) []int {
	arr := make([]int, n)
	arr[0] = 2
	if n > 1 {
		arr[1] = 3
	}
	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr

}

// soal no 4
func reverseNumber(arr []int) []int {
	for i := 1; i < len(arr)/2; i++ {
		temp := arr[len(arr)-i]
		arr[len(arr)-i] = arr[i-1]
		arr[i-1] = temp
	}
	return arr
}

// soal no 5
func isPrimeNumber(num int) bool {
	// 	if num <= 1 {
	// 		return false
	// 	}
	// 	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
	// 		if num%i == 0 {
	// 			return false
	// 		}
	// 	}
	// 	return true

	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true

}

func isPrimes(n int) []string {
	// 	for i := 1; i <= n; i++ {
	// 		isPrimeNumber := isPrimeNumber(i)
	// 		primes[i-1] = fmt.Sprintf("%d=%t", i, isPrimeNumber)
	// 	}
	// 	return primes

	primes := make([]string, n+1)
	for i := 1; i <= n; i++ {
		if isPrimeNumber(i) {
			primes[i] = (strconv.Itoa(i) + " = True")
		} else {
			primes[i] = (strconv.Itoa(i) + " = False")
		}
	}
	return primes
}

func main() {
	// soal no 1
	fmt.Println(evenOrOddNumber(5))

	// soal no 2
	fmt.Println(powNumber(6))

	// soal no 3
	fmt.Println(fibonnaci(5))

	// soal no 4
	fmt.Println(reverseNumber([]int{4, 3, 6, 7, 8, 1}))

	// soal no 5
	fmt.Println(isPrimes(10))

	// n := 50
	// primes := generatePrimes(n)
	// for _, prime := range primes {
	// 	fmt.Println(prime)
	// }
}
