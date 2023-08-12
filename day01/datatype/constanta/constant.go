package main

//global constant variable
const phi = 3.14159

func declareConstant() {
	//constant type
	const price float64 = 20.98 //float64 untuk uang, dan ga bisa di assign/diganti

	//constant untype
	const tax = 11.99

	//multiple constant
	const (
		discount = 0.1 //bisa juga ditulis .01 karena 0 dot bisa tidak ditulis
		phi      = 3.14159
	)

	println(price, tax, discount, phi)
}

func declareIotas() {
	const (
		seq1 = iota
		seq2
		seq3
	)
	var circleArea = phi * 5
	println(circleArea)
	println(seq1, seq2, seq3)
}

func main() {
	declareConstant()
	declareIotas()
}
