package main

import (
	"encoding/json"
	"fmt"
)

type EmployeeDetails struct {
	Email   string `json:"email"`
	Address string `json:"addr"`
	PhoneNo int    `json:"phone_no"`
}

type NamedEmployee struct {
	Name       string          `json:"name"`
	EmpID      int             `json:"emp_id"`
	EmpDetails EmployeeDetails `json:"emp_details"`
}

func main() {
	var namedEmp = NamedEmployee{Name: "Logesh", EmpID: 1,
		EmpDetails: EmployeeDetails{Email: "logesh@my.org", Address: "some addr", PhoneNo: 94389489}}
	// fmt.Println(namedEmp.Email)
	// fmt.Println(namedEmp.Address)
	// fmt.Println(namedEmp.PhoneNo)
	fmt.Println()
	nempJSON, _ := json.MarshalIndent(namedEmp, "", "  ")
	fmt.Println(string(nempJSON))

}

// var namedEmp = NamedEmployee{Name: "Logesh", EmpID: 1, EmpDetails: EmployeeDetails{Email: "logesh@my.org", Address: "some addr", PhoneNo: 94389489}}
// fmt.Println(namedEmp.Email)
// fmt.Println(namedEmp.Address)
// fmt.Println(namedEmp.PhoneNo)

// emp1JSON, _ := json.MarshalIndent(emp1, "", "  ")
// fmt.Println(string(emp1JSON))
