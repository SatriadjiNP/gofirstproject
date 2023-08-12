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

type ProductList []Product

func initProduct() []Product {
	return []Product{
		{"Apple", "Laptop", 1200.98, &Supplier{}},
		{"Samsung", "Gadget", 459.98, &Supplier{}},
		{"Oppo", "Gadget", 230.98, &Supplier{}},
		{"Vivo", "Gadget", 789.98, &Supplier{}},
		{"Xiaomi", "Laptop", 199.98, &Supplier{}},
	}
}

// kumpulan function, implementasi printinfo bisa ditempelkan ke struct product/supplier
// interface header
type Info interface {
	printInfo()
	// untuk menghitung price, berdasarkan category
	totalPrice(tax float64)
}

// implementasi interface, supaya dia bisa dipanggil dengan struct berikut
func (p *Product) printInfo() {
	fmt.Println("Product Info :", p)
}

func (p *ProductList) totalPrice(tax float64) map[string]float64 {
	totals := make(map[string]float64)
	for _, v := range *p {
		// category key dari list
		totals[v.category] = totals[v.category] + (v.price * tax)
	}
	return totals
}

func (s *Supplier) printinfo() {
	fmt.Println("Supplier Info :", s)
}

func main() {
	products := ProductList(initProduct())

	products[0].printInfo()

	s := Supplier{id: 1, supplierName: "ABC"}
	s.printinfo()
	fmt.Println()

	products.totalPrice(0.5)
	fmt.Println(products.totalPrice(0.5))
}
