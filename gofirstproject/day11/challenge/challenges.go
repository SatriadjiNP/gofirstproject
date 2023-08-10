package main

import "fmt"

type Employee struct {
	empId       int
	firstName   string
	lastName    string
	role        string
	basicSalary float64
	tunjangan   *Tunjangan
	totalSalary float64
}

type Tunjangan struct {
	bonus      float64
	commission float64
	lembur     float64
	makan      float64
}

type EmployeeBasic struct {
	employee []Employee
}

// func generateId(empId *Employee) int {
// 	var id = 119
// 	id++
// 	return id
// }

// constructor
func employeeData(empId int, firstName, lastName, role string, basicSalary float64, tunjangan *Tunjangan, totalSalary float64) *Employee {
	return &Employee{empId, firstName, lastName, role, basicSalary, tunjangan, totalSalary}
}

func (print *Employee) Print() {
	fmt.Printf("EmpID: %d\n", print.empId)
	fmt.Printf("Name: %s %s\n", print.firstName, print.lastName)
	fmt.Printf("Role: %s\n", print.role)
	fmt.Printf("Basic Salary: %.2f\n", print.basicSalary)
	fmt.Printf("Tunjangan Bonus: %.2f\n", print.tunjangan.bonus)
	fmt.Printf("Tunjangan Commission: %.2f\n", print.tunjangan.commission)
	fmt.Printf("Tunjangan Lembur: %.2f\n", print.tunjangan.lembur)
	fmt.Printf("Tunjangan Makan: %.2f\n", print.tunjangan.makan)
	fmt.Printf("Total Salary: %.2f\n\n", print.totalSalary)

}

// method, total salary
func (empl *Employee) calcTotalSalary(totalSalary *Employee) float64 {
	if totalSalary.role == "Programmer" {
		totalSalary.totalSalary = totalSalary.basicSalary + empl.tunjangan.lembur
	}
	if totalSalary.role == "Sales" {
		totalSalary.totalSalary = totalSalary.basicSalary + empl.tunjangan.bonus + empl.tunjangan.commission
	}
	if totalSalary.role == "QA" {
		totalSalary.totalSalary = totalSalary.basicSalary + empl.tunjangan.makan
	}
	return empl.totalSalary
}

// method, basic salary all
func (empl *EmployeeBasic) calcTotalBasicSalary() float64 {
	var basicAll float64
	for _, v := range empl.employee {
		basicAll += v.basicSalary

	}
	return basicAll
}

// method, total salary all
func (empl *EmployeeBasic) calcTotalSalaryAll() float64 {
	var salaryAll float64
	for _, v := range empl.employee {
		salaryAll += v.totalSalary

	}
	return salaryAll
}

// method, total salary role
func (empl *EmployeeBasic) calcTotalSalaryRole(role string) float64 {
	var salaryRole float64
	// var salarySales float64
	// var salaryQA float64
	for _, v := range empl.employee {
		if v.role == role {
			salaryRole += v.totalSalary
		}

	}
	return salaryRole
}

func main() {
	// // Menghasilkan ID baru
	// id := generateId()

	// // Menampilkan ID
	// fmt.Println("empId:", id)
	tunjangan1 := &Tunjangan{0, 0, 500000, 0}
	tunjangan2 := &Tunjangan{0, 0, 500000, 0}
	tunjangan3 := &Tunjangan{0, 0, 500000, 0}
	tunjangan4 := &Tunjangan{500000, 200000, 0, 0}
	tunjangan5 := &Tunjangan{350000, 250000, 0, 0}
	tunjangan6 := &Tunjangan{0, 0, 0, 10000}
	tunjangan7 := &Tunjangan{0, 0, 0, 10000}

	emp1 := employeeData(120, "Anton", "Pramata", "Programmer", 6_000_000, tunjangan1, 0)
	emp2 := employeeData(121, "Budi", "Junaedi", "Programmer", 6_000_000, tunjangan2, 0)
	emp3 := employeeData(122, "Charlie", "Van Dijk", "Programmer", 6_000_000, tunjangan3, 0)
	emp4 := employeeData(123, "Dian", "Permana", "Sales", 3_000_000, tunjangan4, 0)
	emp5 := employeeData(124, "Fatur", "Rohman", "Sales", 3_000_000, tunjangan5, 0)
	emp6 := employeeData(125, "Ellise", "Toon", "QA", 4_500_000, tunjangan6, 0)
	emp7 := employeeData(126, "Gita", "Gutawa", "QA", 4_500_000, tunjangan7, 0)

	// empPlus1 := 0.0
	// empPlus1 += emp1.basicSalary + tunjangan1.lembur
	// fmt.Println(emp1.calcTotalSalary(empPlus1))
	// emp1.totalSalary = emp1.basicSalary + tunjangan1.lembur
	emp1.calcTotalSalary(emp1)
	emp2.calcTotalSalary(emp2)
	emp3.calcTotalSalary(emp3)

	emp4.calcTotalSalary(emp4)
	emp5.calcTotalSalary(emp5)

	emp6.calcTotalSalary(emp6)
	emp7.calcTotalSalary(emp7)

	// emp4.calcTotalSalary(emp4)
	// emp4.calcTotalSalary(emp4)
	emp := []Employee{*emp1, *emp2, *emp3, *emp4, *emp5, *emp6, *emp7}
	for _, value := range emp {
		value.Print()
	}

	fmt.Println("Total Employee : ", len(emp))

	fmt.Println("Total Role Employee")
	roles := make(map[string]int)
	for _, value := range emp {
		roles[value.role]++
	}
	for role, count := range roles {
		fmt.Printf("%s : %d\n", role, count)
	}

	totalBasic := EmployeeBasic{employee: emp}
	totalBasicSalary := totalBasic.calcTotalBasicSalary()

	fmt.Printf("Total Basic All : %.2f", totalBasicSalary)
	fmt.Println()

	totalSalary := EmployeeBasic{employee: emp}
	totalSalaryAll := totalSalary.calcTotalSalaryAll()

	fmt.Printf("Total Salary All : %.2f", totalSalaryAll)
	fmt.Println()

	totalSalaryProg := EmployeeBasic{employee: emp}
	totalSalarySal := EmployeeBasic{employee: emp}
	totalSalaryqa := EmployeeBasic{employee: emp}

	totalSalaryAllProg := totalSalaryProg.calcTotalSalaryRole("Programmer")
	totalSalaryAllSal := totalSalarySal.calcTotalSalaryRole("Sales")
	totalSalaryAllQA := totalSalaryqa.calcTotalSalaryRole("QA")

	fmt.Printf("Total Salary Programmer : %.2f\n", totalSalaryAllProg)
	fmt.Printf("Total Salary Sales : %.2f\n", totalSalaryAllSal)
	fmt.Printf("Total Salary QA : %.2f", totalSalaryAllQA)
	fmt.Println()

}
