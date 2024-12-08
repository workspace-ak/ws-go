package main

import "fmt"

type Address struct {
	Street string
	City   string
}

type Employee struct {
	Name     string
	Age      int
	IsRemote bool
	Address
}

func (e *Employee) editName(newName string) {
	e.Name = newName
}

func (a Address) printAddress() {
	fmt.Printf("[-] %s, %s\n", a.Street, a.City)
}

func main() {
	adrs1 := Address{
		Street: "B-52, Rohit Apartment",
		City:   "Delhi",
	}

	employee1 := Employee{
		Name:     "Bob",
		Age:      32,
		IsRemote: true,
		Address:  adrs1,
	}

	fmt.Println("[a] Employee Name: ", employee1.Name)
	fmt.Println("[a] Remote: ", employee1.IsRemote)

	// Anonymous struct
	job := struct {
		title  string
		salary int
	}{
		title:  "Data Scientist",
		salary: 25000,
	}

	fmt.Println("[b] Salary: ", job.salary)
	fmt.Println("[b] Title: ", job.title)

	employeePtr := &employee1
	employeePtr.Age = 31
	fmt.Println("[c] Employee Age: ", employee1.Age)

	employee1.editName("Stevens")
	fmt.Println("[d] Employee Name: ", employee1.Name)

	fmt.Println("[e] Employee Address: ", employee1.Street, employee1.City)
	employee1.printAddress()
}
