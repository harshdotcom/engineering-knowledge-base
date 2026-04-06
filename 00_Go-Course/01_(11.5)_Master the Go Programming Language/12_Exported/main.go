package main

import (
	"fmt"
	"myApp/staff"
)

var employees = []staff.Employee{
	{FirstName: "John", LastName: "Smith", Salary: 3000, FullTime: true},
	{FirstName: "Harsh", LastName: "Smith", Salary: 3000, FullTime: true},
	{FirstName: "Sree", LastName: "Smith", Salary: 3000, FullTime: true},
	{FirstName: "Rahul", LastName: "Smith", Salary: 3000, FullTime: true},
	{FirstName: "Shyam", LastName: "Smith", Salary: 3000, FullTime: true},
	{FirstName: "Kalyani", LastName: "Smith", Salary: 3000, FullTime: true},
}

func main() {
	myStaff := staff.Office{
		AllStaff: employees,
	}

	myStaff.All()

	fmt.Println(myStaff.All())
	fmt.Println(myStaff.OverPaid())

}
