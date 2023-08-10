package main

import "fmt"

// tujuan pointer adalah untuk menghindari copy data/value yang terlalu besar

// & get the address of the variable
// * get the value of pointer

// copy by value
func copyByValue() {
	first := 100
	second := first

	first++

	fmt.Println(first, "address 1st : ", &first)
	fmt.Println(second, "address 2nd : ", &second)
}

// copy by pointer
func copyByPointer() {
	first := 100
	second := &first

	first++

	fmt.Println(first, "address 1st : ", &first)
	fmt.Println(second, "address 2nd : ", &second)
	fmt.Println(*second, "value of pointer : ", second)
}

func chainPointer() {
	first := 100
	second := &first
	third := &second

	first++
	*second++
	**third++

	fmt.Println(first, "address 1st : ", &first)
	fmt.Println(second, "address 2nd : ", &second)
	fmt.Println(third, "address 3rd : ", &third)

}

func pointerZeroValue() {
	first := 100
	// jika kita coba akses, muncul panic error :invalid memory address
	var second *int
	second = &first

	fmt.Println("First : ", first)
	fmt.Println("Second : ", *second)
}

func slicePointer() {
	// bila ingin mengakses address setiap elemen, harus gunakan looping
	fruit := []string{"Apel", "Mangga", "Pisang", "Melon", "Nanas", "Durian", "Semangka", "Jeruk"}

	for key := range fruit {
		// %d interger, %p desimal/hexa, %s string
		fmt.Printf("[%d] %p %s\n", key, &fruit[key], fruit[key])

	}
}

func slicePointerAddr() {
	// bila ingin mengakses address setiap elemen, harus gunakan looping
	fruit := []string{"Apel", "Mangga", "Pisang", "Melon"}
	fruit2 := []string{"Nanas", "Durian", "Semangka", "Jeruk"}
	for key := range fruit {
		// %d interger, %p desimal/hexa, %s string
		fmt.Printf("[%d] %p %s\n", key, &fruit[key], fruit[key])
	}

	fruit = append(fruit, fruit2...)
	// bila append di pointer, maka akan membuat array baru dikarenakan address berubah menjadi baru juga
	for key := range fruit {
		// %d interger, %p desimal/hexa, %s string
		fmt.Printf("[%d] %p %s\n", key, &fruit[key], fruit[key])
	}
}

func main() {
	// copyByValue()
	// copyByPointer()
	// chainPointer()
	// pointerZeroValue()
	// slicePointer()
	slicePointerAddr()
}
