package main

import (
	"fmt"
)

/*
Persoalan kasus
-------------------------------------------------------------------------------------------------------------------------------------------
1           Anton       Pratama     Programmer  Rp. 6000000       Rp. 0           Rp. 0         Rp. 500000          Rp. 0       Rp. 6500000
2           Budi        Junaedi     Programmer  Rp. 6000000       Rp. 0           Rp. 0         Rp. 500000          Rp. 0       Rp. 6500000
3           Charlie     Van Dijk    Programmer  Rp. 6000000       Rp. 0           Rp. 0         Rp. 500000          Rp. 0       Rp. 6500000
4           Dian        Permana     Sales       Rp. 3000000       Rp. 500000      Rp. 200000    Rp. 0               Rp. 0       Rp. 3700000
5           Fatur       Rohman      Sales       Rp. 3000000       Rp. 350000      Rp. 250000    Rp. 0               Rp. 0       Rp. 3600000
6           Ellise      Toon        QA          Rp. 4500000       Rp. 0           Rp. 0         Rp. 0               Rp. 10000   Rp. 4510000
7           Gita        Gutawa      QA          Rp. 4500000       Rp. 0           Rp. 0         Rp. 0               Rp. 10000   Rp. 4510000
-------------------------------------------------------------------------------------------------------------------------------------------
*/

// nilai konstan, lembur dan makan
const (
	LEMBUR = 500_000
	MAKAN  = 10_000
)

// untuk menggunakan pointer, tergantung ketika data sama dan bisa untuk dishare
// 1. create struct
// struct berfungsi untuk hold data, dan pastikan data mana yang mau dihold
type Employee struct {
	empId                     int
	firstName, lastName, role string
	basicSalary               float64
	tunjangan                 Allowances
	totalSalary               float64
}

type Allowances struct {
	sales      TunjSales
	programmer TunjProgram
	qa         TunjQA
}

type TunjSales struct {
	bonus, commission float64
}

type TunjProgram struct {
	lembur float64
}

type TunjQA struct {
	makan float64
}

type EmployeeList []Employee

var countId = 119

// 2. constructor, hanya bisa 1 record
func newEmployee(firstName, lastName, role string, basicSalary float64, allowances Allowances) *Employee {
	countId++
	return &Employee{
		countId,
		firstName, lastName, role,
		basicSalary, allowances, setTotalSalary(basicSalary, allowances)}
}

// 3. create initEmployee
func initEmployee() EmployeeList {
	return EmployeeList{
		*newEmployee("Anton", "Pramata", "Programmer", 6_000_000,
			calTunjangan("Programmer", Allowances{
				sales:      TunjSales{},
				programmer: TunjProgram{LEMBUR},
				qa:         TunjQA{},
			})),
		*newEmployee("Budi", "Junaedi", "Programmer", 6_000_000,
			calTunjangan("Programmer", Allowances{
				sales:      TunjSales{},
				programmer: TunjProgram{LEMBUR},
				qa:         TunjQA{},
			})),
		*newEmployee("Charlie", "Van Dijk", "Programmer", 6_000_000,
			calTunjangan("Programmer", Allowances{
				sales:      TunjSales{},
				programmer: TunjProgram{LEMBUR},
				qa:         TunjQA{},
			})),
		*newEmployee("Dian", "Permana", "Sales", 3_000_000,
			calTunjangan("Sales", Allowances{
				sales:      TunjSales{bonus: 500_000, commission: 200_000},
				programmer: TunjProgram{},
				qa:         TunjQA{},
			})),
		*newEmployee("Fatur", "Rohman", "Sales", 3_000_000,
			calTunjangan("Sales", Allowances{
				sales:      TunjSales{bonus: 350_000, commission: 250_000},
				programmer: TunjProgram{},
				qa:         TunjQA{},
			})),
		*newEmployee("Ellise", "Toon", "QA", 4_500_000,
			calTunjangan("QA", Allowances{
				sales:      TunjSales{},
				programmer: TunjProgram{},
				qa:         TunjQA{MAKAN},
			})),
		*newEmployee("Gita", "Gutama", "QA", 4_500_000,
			calTunjangan("QA", Allowances{
				sales:      TunjSales{},
				programmer: TunjProgram{},
				qa:         TunjQA{MAKAN},
			})),
	}
	// return EmployeeList{*emp1, *emp2, *emp3, *emp4, *emp5, *emp6, *emp7}
}

// function calculate tunjangan
func calTunjangan(role string, allow Allowances) Allowances {
	tunjangan := Allowances{}
	switch role {
	case "Sales":
		tunjangan.sales.bonus = allow.sales.bonus
		tunjangan.sales.commission = allow.sales.commission
	case "Programmer":
		tunjangan.programmer.lembur = LEMBUR
	case "QA":
		tunjangan.qa.makan = MAKAN
	}
	return tunjangan
}

// func untuk menghitung total salary, baru dipanggil di method atas
func setTotalSalary(basicSalary float64, allow Allowances) float64 {
	return basicSalary + allow.sales.bonus + allow.sales.commission + allow.programmer.lembur + allow.qa.makan
}

// 4. create interface
type summary interface {
	setBasicSalary(salary float64)
}

func (emp *Employee) setBasicSalary(salary float64) {
	emp.basicSalary = salary
	emp.totalSalary = emp.basicSalary + emp.tunjangan.programmer.lembur +
		emp.tunjangan.sales.bonus + emp.tunjangan.sales.commission + emp.tunjangan.qa.makan
}

func main() {
	employeeList := initEmployee()
	employeeList[1].setBasicSalary(7_000_000)

	for _, v := range employeeList {
		fmt.Println(v)
	}

}
