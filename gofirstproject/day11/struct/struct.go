package main

import "fmt"

// stack memory data kecil
// keep memory data besar, dan ada garbage memory

type Product struct {
	name, category string
	price          float64
}

type OrderProduct struct {
	product    Product // embedded fields
	totalOrder int     // regular fields
}

func initProduct() {
	prod := Product{
		name:     "Laptop",
		category: "Electronic",
		price:    250,
	}
	// before update
	fmt.Println(prod)

	// update value
	prod.name = "Dell Laptop"
	prod.price = 350.98
	fmt.Println(prod)

	// default zero value
	var tv Product
	tv.name = "Sharp"

	meja := Product{}
	fmt.Println(tv, meja)
}

func initProductPointer() {
	// built-in new function, result pinter struct
	var iphone = new(Product)
	iphone.name = "Iphone"
	iphone.category = "Gadget"
	iphone.price = 1500

	fmt.Println(iphone)

	// using &pointer
	var oppo = &Product{}
	oppo.name = "Oppo"
	oppo.category = "Gadget"
	oppo.price = 1000

	fmt.Println(oppo)
}

// memanggil type order
func initOrderProduct() {
	myOrder := OrderProduct{
		product:    Product{"Samsung", "Gadget", 1258},
		totalOrder: 1,
	}
	fmt.Println(myOrder)
}

func compareStruct() {

	// kalo memakai ini otomatis addressnya beda dan hasil akan false
	// var oppo = &Product{}
	// oppo.name = "Oppo"
	// oppo.category = "Gadget"
	// oppo.price = 876

	// var iphone = new(Product)
	// iphone.name = "Iphone"
	// iphone.category = "Gadget"
	// iphone.price = 1500

	oppo := Product{
		name: "Oppo", category: "Gadget", price: 855,
	}
	iphone := Product{
		name: "Oppo", category: "Gadget", price: 855,
	}
	samsung := Product{
		name: "Samsung", category: "Gadget", price: 1245,
	}
	xiaomi := Product{
		name: "Xiaomi", category: "Gadget", price: 256,
	}

	fmt.Println("Compare : ", oppo == iphone)
	fmt.Println("Compare : ", samsung == xiaomi)

	printStruct(oppo)
}

// anonymous struct
func printStruct(val struct {
	name, category string
	price          float64
}) {
	fmt.Println("Name :", val.name, "\nCategory :", val.category, "\nPrice :", val.price)
}

func findAllProduct() []Product {
	oppo := Product{
		name: "Oppo", category: "Gadget", price: 755,
	}
	iphone := Product{
		name: "Iphone", category: "Gadget", price: 1882,
	}
	xiaomi := Product{
		name: "Xiaomi", category: "Gadget", price: 256,
	}

	products := []Product{oppo, iphone, xiaomi}

	// memakai cara append
	products = append(products, Product{name: "Samsung", category: "Gadget", price: 1225})

	return products

}

// misalkan data unique samsung, oppo, iphone
func findProductByName(products []Product, search string) Product {
	for _, product := range products {
		if product.name == search {
			return product
		}
	}
	// Jika tidak ditemukan, mengembalikan produk kosong
	return Product{}
}

// CONTOH DARI CHATGPT
// Struct Person
type Person struct {
	Name string
	Age  int
}

// Fungsi untuk mengubah nama menggunakan pointer
func changeName(p *Person, newName string) {
	p.Name = newName
}

func main() {
	// initProduct()
	// initProductPointer()
	// initOrderProduct()
	// compareStruct()
	// findAllProduct()
	fmt.Println(findProductByName(findAllProduct(), "Samsung"))
	fmt.Println(findProductByName(findAllProduct(), "Xiaomi"))

	// // Membuat objek Person dengan pointer
	// p := &Person{Name: "John Doe", Age: 30}

	// // Mengakses dan mengubah nilai menggunakan pointer
	// fmt.Println("Nama awal:", p.Name) // Output: Nama awal: John Doe
	// changeName(p, "Jane Smith")
	// fmt.Println("Nama setelah diubah:", p.Name) // Output: Nama setelah diubah: Jane Smith

}
