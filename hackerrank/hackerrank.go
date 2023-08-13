package main

import (
	"fmt"
	"strconv"
)

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

func plusMinus(arr []int32) {
	// Write your code here
	size := len(arr)
	var num1, num2, zero int

	for _, v := range arr {
		switch {
		case v > 0:
			num1++
		case v < 0:
			num2++
		default:
			zero++
		}
	}
	ratio1 := float64(num1) / float64(size)
	ratio2 := float64(num2) / float64(size)
	ratioZero := float64(zero) / float64(size)
	fmt.Printf("%f\n", ratio1)
	fmt.Printf("%f\n", ratio2)
	fmt.Printf("%f\n", ratioZero)
}

func staircase(n int32) {
	for i := int32(1); i <= n; i++ {
		for j := int32(1); j <= n-i; j++ {
			fmt.Print(" ")
		}
		for k := int32(1); k <= i; k++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}

func miniMaxSum(arr []int32) {
	var minSum, maxSum, totalSum int64

	for _, num := range arr {
		totalSum += int64(num)
	}

	minSum = totalSum - int64(arr[0])
	maxSum = totalSum - int64(arr[0])

	for _, num := range arr {
		sumWithoutNum := totalSum - int64(num)
		if sumWithoutNum < minSum {
			minSum = sumWithoutNum
		}
		if sumWithoutNum > maxSum {
			maxSum = sumWithoutNum
		}
	}

	fmt.Println(minSum, maxSum)
}

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	var tallestHeight int32
	tallestCount := int32(0)

	for _, height := range candles {
		if height > tallestHeight {
			tallestHeight = height
			tallestCount = 1
		} else if height == tallestHeight {
			tallestCount++
		}
	}

	return tallestCount
}

func timeConversion(s string) string {
	// Parse the input time string
	hour, _ := strconv.Atoi(s[:2])
	period := s[len(s)-2:]

	// Handle the special cases of 12AM and 12PM
	if period == "AM" && hour == 12 {
		hour = 0
	} else if period == "PM" && hour != 12 {
		hour += 12
	}

	// Convert hour back to string format
	hourStr := strconv.Itoa(hour)
	if hour < 10 {
		hourStr = "0" + hourStr
	}

	return hourStr + s[2:len(s)-2]
}

func gradingStudents(grades []int32) []int32 {
	var roundedGrades []int32

	for _, grade := range grades {
		if grade < 38 {
			roundedGrades = append(roundedGrades, grade)
		} else {
			nextMultipleOf5 := (grade/5 + 1) * 5
			if nextMultipleOf5-grade < 3 {
				roundedGrades = append(roundedGrades, nextMultipleOf5)
			} else {
				roundedGrades = append(roundedGrades, grade)
			}
		}
	}

	return roundedGrades
}

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	appleCount := 0
	orangeCount := 0

	for _, apple := range apples {
		if a+apple >= s && a+apple <= t {
			appleCount++
		}
	}

	for _, orange := range oranges {
		if b+orange >= s && b+orange <= t {
			orangeCount++
		}
	}

	fmt.Println(appleCount)
	fmt.Println(orangeCount)
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

	println()

	fmt.Println("Soal plusMinus")
	numbers := []int32{-4, 3, -9, 0, 4, 1}
	plusMinus(numbers)

	println()

	fmt.Println("Soal Staircase")
	staircase(6)

	println()

	fmt.Println("Soal MiniMaxSum")
	num := []int32{1, 2, 3, 4, 5}
	miniMaxSum(num)

	println()

	fmt.Println("Soal birthdayCakeCandles")
	candles := []int32{3, 2, 1, 3}
	fmt.Println(birthdayCakeCandles(candles))

	println()

	fmt.Println("Soal timeConversion")
	fmt.Println(timeConversion("07:05:45PM"))

	println()

	fmt.Println("Soal gradingStudents")
	fmt.Println(gradingStudents([]int32{73, 67, 38, 33}))

	println()

	fmt.Println("Soal countApplesAndOranges")
	s := int32(7)
	t := int32(11)
	a := int32(5)
	b := int32(15)
	apples := []int32{-2, 2, 1}
	oranges := []int32{5, -6}
	countApplesAndOranges(s, t, a, b, apples, oranges)
}
