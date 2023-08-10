package main

import "fmt"

type Person struct {
	name string
	age  string
}

// immutable : tidak bisa diubah data originnya
func immutable(i int, s string, p Person) {
	i = i * 5
	s = "Bootcamp Go"
	p.name = "Miracle"
}

// mutable, copy by pointer
// kapan memakai pointer = ketika data originnya mau diubah
func mutable(i *int, s *string, p *Person) {
	*i = *i * 5
	*s = "Bootcamp Go"
	p.name = "Miracle"
}

func modMap(m map[int]string) {
	m[2] = "Second"
	m[3] = "Third"
	delete(m, 1)
}

func main() {
	// sample immutable, copy value
	// meskipun dijalankan func, ketika dicetak tetap data originnya, copy value
	p := Person{}
	i := 10
	s := "Java"
	immutable(i, s, p)
	fmt.Println(i, s, p)

	// call mutable
	mutable(&i, &s, &p)
	fmt.Println(i, s, p)

	// modifikasi map
	m := map[int]string{
		1: "First",
		2: "Dua",
	}
	modMap(m)
	fmt.Println(m)
}
