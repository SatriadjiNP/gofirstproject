package main

import "fmt"

func initMap() {
	// store nil
	emps := map[string]int{}
	fmt.Println(emps)

	// using make
	products := make(map[string]float64)
	products["Apple"] = 90.00
	products["Orange"] = 100.00
	products["Banana"] = 110.98
	// bila key sama, maka akan mengupdate value yang terupdate
	products["Apple"] = 100.00
	fmt.Println(products)

	fmt.Println("Price Fruits : ", products["Apple"])

	// literal value
	students := map[string]float64{
		"A": 100.00,
		"B": 90.00,
		"C": 96.80,
	}
	for key, value := range students {
		fmt.Println("Key : ", key, "Value : ", value)
	}
}

func deleteMap() {
	students := map[string]float64{
		"A": 100.00,
		"B": 90.00,
		"C": 96.80,
	}
	delete(students, "B")

	// check key value A
	if value, ok := students["A"]; ok {
		fmt.Println("Key A : already exist -", value)
		delete(students, "A")
		fmt.Println(value, "Has been deleted")
	} else {
		fmt.Println("No key store ready")
	}

	fmt.Println(students)
}

// filtering test yang nomor 7 pake cara ini ggwp
func totalFruits(fruit []string, vege []string) map[string]int {
	mergeArray := []string{}
	// variadic operator
	mergeArray = append(mergeArray, fruit...) // yang disisipkan berbentuk array, maka ...
	mergeArray = append(mergeArray, vege...)

	// contoh return map
	mapFruit := map[string]int{}

	// declare mapFruit to store key, value
	for key := range mergeArray {
		mapFruit[mergeArray[key]]++
	}
	return mapFruit

}

// map tidak bisa mensorting
func diffFruit(fruit []string, vege []string) ([]string, []string) {
	mergeArray := []string{}
	// variadic operator
	mergeArray = append(mergeArray, fruit...) // yang disisipkan berbentuk array, maka ...
	mergeArray = append(mergeArray, vege...)

	// contoh return map
	mapFruit := map[string]int{}

	// declare mapFruit to store key, value
	// kalo ini langsung sorting
	for key := range mergeArray {
		mapFruit[mergeArray[key]]++
	}

	// declare new variable same array
	sameArray := []string{}
	diffArray := []string{}

	for key, value := range mapFruit {
		if value > 1 {
			sameArray = append(sameArray, key)
		} else {
			diffArray = append(diffArray, key)
		}
	}
	return sameArray, diffArray
}

func main() {
	initMap()
	deleteMap()
	result := totalFruits(
		[]string{"Apple", "Melon", "Watermelon", "Banana"},
		[]string{"Tomato", "Apple", "Melon", "Banana", "Spring Onions"})
	fmt.Println(result)

	sameArray, diffArray := diffFruit(
		[]string{"Apple", "Melon", "Watermelon", "Banana"},
		[]string{"Tomato", "Apple", "Melon", "Banana", "Spring Onions"})
	fmt.Println("Same : ", sameArray)
	fmt.Println("Diff : ", diffArray)
}
