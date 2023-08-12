package main

import "fmt"

type ListEmployees struct {
	EmpId                            int
	FirstName                        string
	LastName                         string
	Role                             string
	BasicSalary                      float64
	Bonus, Commission, Lembur, Makan float64
	TotalSalary                      float64
}

type Employee []ListEmployees

var lastEmpId int = 119

//constructor
func dataEmployee() []ListEmployees {
	return []ListEmployees{
		{120, "Anton", "Pratama", "Programmer", 6000000, 0, 0, 500000, 0, 0},
		{121, "Budi", "Junaedi", "Programmer", 6000000, 0, 0, 500000, 0, 0},
		{122, "Charlie", "Van Dijk", "Programmer", 6000000, 0, 0, 500000, 0, 0},
		{123, "Dian", "Permana", "Sales", 3000000, 500000, 200000, 0, 0, 0},
		{124, "Fatur", "Rohman", "Sales", 3000000, 350000, 250000, 0, 0, 0},
		{125, "Ellise", "Toon", "QA", 4500000, 0, 0, 0, 10000, 0},
		{126, "Gita", "Gutawa", "QA", 4500000, 0, 0, 0, 10000, 0},
	}

}

type Info interface {
	totalSalarySum()
	printInfo()
	totalBasicSalaryAll()
	totalSalaryAll()
	totalEmployeeRole(role string)
}

// implementasi interface
func (print *ListEmployees) printInfo() {
	fmt.Printf("EmpID: %d\n", print.EmpId)
	fmt.Printf("Name: %s %s\n", print.FirstName, print.LastName)
	fmt.Printf("Role: %s\n", print.Role)
	fmt.Printf("Basic Salary: %.2f\n", print.BasicSalary)
	fmt.Printf("Tunjangan Bonus: %.2f\n", print.Bonus)
	fmt.Printf("Tunjangan Commission: %.2f\n", print.Commission)
	fmt.Printf("Tunjangan Lembur: %.2f\n", print.Lembur)
	fmt.Printf("Tunjangan Makan: %.2f\n", print.Makan)
	fmt.Printf("Total Salary: %.2f\n\n", print.TotalSalary)
}

func (empl *ListEmployees) totalSalarySum() {
	switch empl.Role {
	case "Programmer":
		empl.TotalSalary = empl.BasicSalary + empl.Lembur
	case "Sales":
		empl.TotalSalary = empl.BasicSalary + empl.Bonus + empl.Commission
	case "QA":
		empl.TotalSalary = empl.BasicSalary + empl.Makan
	default:
		empl.TotalSalary = empl.BasicSalary
	}
}

func (empl *Employee) totalBasicSalaryAll() float64 {
	var basicAll float64
	for _, v := range *empl {
		basicAll += v.BasicSalary

	}
	return basicAll
}

func (empl *Employee) totalSalaryAll() float64 {
	var salaryAll float64
	for _, v := range *empl {
		salaryAll += v.TotalSalary

	}
	return salaryAll
}

func (empl *Employee) totalEmployeeRole(role string) float64 {
	var salaryRole float64
	for _, v := range *empl {
		if v.Role == role {
			salaryRole += v.TotalSalary
		}

	}
	return salaryRole
}

func main() {
	emp := Employee(dataEmployee())
	for _, value := range emp {
		value.totalSalarySum()
		value.printInfo()
	}

	fmt.Println("Total Employee : ", len(emp))

	fmt.Println("Total Role Employee")
	roles := make(map[string]int)
	for _, value := range emp {
		roles[value.Role]++
	}
	for role, count := range roles {
		fmt.Printf("%s : %d\n", role, count)
	}
	fmt.Println()

	totalBasic := emp.totalBasicSalaryAll()
	fmt.Printf("Total Basic Salary : %.2f", totalBasic)
	fmt.Println()

	totalAllSalary := emp.totalSalaryAll()
	fmt.Printf("Total Salary All : %.2f", totalAllSalary)
	fmt.Println()

	totalEmpRole := emp.totalEmployeeRole("Programmer")
	fmt.Println("Total Salary Programmer :", totalEmpRole)

}
