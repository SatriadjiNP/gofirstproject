package main

func divisonModulus() {
	number := 15
	bagi := number / 3
	sisa := number % 3
	println(bagi, sisa)
}

func comparison() {
	first := 100
	const second = 200
	equal := first == second
	notEqual := first != second
	lessThan := first < second
	lessThanOrEqual := first <= second
	greater := first > second
	greaterThanOrEqual := first >= second
	println("First : ", first)
	println("Equal : ", equal)
	println("Less Than : ", lessThan)
	println("Not Equal : ", notEqual)
	println("Less Than Or Equal : ", lessThanOrEqual)
	println("Greater : ", greater)
	println("Greater Than Or Equal : ", greaterThanOrEqual)

}

func main() {
	divisonModulus()
	comparison()
}
