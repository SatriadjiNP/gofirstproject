package main

import "fmt"

type Product struct {
	name, category string
	price          float64
	supplier       *Supplier
}

type Supplier struct {
	id           int
	supplierName string
}

func copyByValue() {
	p1 := Product{name: "Nokia", category: "Mobile", price: 123}

	p2 := p1
	p2.name = "Oppo"

	fmt.Println("p1 :", p1)
	fmt.Println("p2 :", p2)

}

func copyByPointer() {
	p1 := Product{name: "Nokia", category: "Mobile", price: 123}

	// meskipun p2.name yang diubah, tetapi kalo pointer tetap sama dengan p1
	p2 := &p1
	p2.name = "Oppo"

	fmt.Println("p1 :", p1)
	fmt.Println("p2 :", *p2)

}

func calcTax(product *Product) {
	if product.price > 100 {
		product.price = product.price * 0.2
	}
	fmt.Println("Total Price + Tax :", product)
}

// contoh memakai return
func calcTaxReturn(product *Product) *Product {
	if product.price > 100 {
		product.price = product.price * 0.2
	}

	return product
}

// function constructor
func newProduct(name, category string, price float64, supplier *Supplier) *Product {
	// discount := 10
	return &Product{name, category, price - (price * 0.1), supplier}
}

// menggunakan method, harus dibuat diawal function, dinamakan receiver type
func (product *Product) calcDiscount(discount float64) float64 {
	if product.price > 100 {
		return product.price - (product.price * discount)
	}
	return product.price
}

func main() {
	copyByValue()
	fmt.Println()
	copyByPointer()
	fmt.Println()

	prod := Product{name: "Nokia", category: "Mobile", price: 123}
	calcTax(&prod)
	fmt.Println()

	prod1 := Product{name: "Esia", category: "Mobile", price: 123}
	fmt.Println(calcTaxReturn(&prod1))
	fmt.Println()

	// initial constructor
	supplier := &Supplier{1, "Agent Car"}

	car1 := newProduct("BMW", "Car", 3000, supplier)
	car2 := newProduct("Lambo", "Car", 2000, supplier)

	car1.price = car1.calcDiscount(0.5)
	car2.price = car2.calcDiscount(0.3)

	car1.supplier.supplierName = "Agent Motor"
	// value
	product := []Product{*car1, *car2}
	// address
	// product := []*Product{car1, car2}

	fmt.Println(product)
}
