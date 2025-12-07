package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) printfPerson() {
	println("Name:", p.Name)
	println("Age:", p.Age)
}

type Employee struct {
	Person
	EmployeeID int
}

func (e *Employee) printEmployee() {
	e.printfPerson()
	println("EmployeeID:", e.EmployeeID)
}

func (e *Employee) printEmployee11() {
	fmt.Printf("e.Name: %v\n", e.Name)
	fmt.Printf("e.Age: %v\n", e.Age)
	fmt.Printf("e.EmployeeID: %v\n", e.EmployeeID)
}

func main() {
	emp := Employee{
		Person:     Person{Name: "Bingtang", Age: 30},
		EmployeeID: 12345,
	}
	emp.printEmployee()
	emp.printEmployee11()

}
