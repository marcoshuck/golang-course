package main

import (
	"fmt"
	"time"
)

type EmployeeIdentification interface {
	WhoAmI() string
}

type Person struct {
	FirstName string
	LastName  string
	Birthday  time.Time
}

type Role struct {
	Name  string
	Power uint
}

type Employee struct {
	Person Person
	Role   Role
}

//func (e Employee) WhoAmI() string {
//	return fmt.Sprintf("%s %s", e.Person.FirstName, e.Person.LastName)
//}

type SuperEmployee struct {
	Employee
	SuperPower
}

type SuperPower struct {
	Name string
}

func main() {
	e := Employee{
		Person: Person{
			FirstName: "Marcos",
			LastName:  "Huck",
			Birthday:  time.Now(),
		},
		Role: Role{
			Name:  "Admin",
			Power: 1337,
		},
	}

	se := SuperEmployee{
		Employee: e,
		SuperPower: SuperPower{
			Name: "Great teacher!",
		},
	}

	// Cannot use 'se' (type SuperEmployee) as the type EmployeeIdentification
	// Type does not implement 'EmployeeIdentification' as some methods are missing: WhoAmI() string
	PrintWhoAmI(se)

	fmt.Printf("%+v\n", se)
}

func PrintWhoAmI(em EmployeeIdentification) {
	fmt.Println("Who am I?", em.WhoAmI())
}
