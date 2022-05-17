package main

import (
	"fmt"
)

type EmployeeDetails struct {
	Email   string `json:"email"`
	Address string `json:"addr"`
	PhoneNo int    `json:"phone_no"`
}

type Employee struct {
	Name            string `json:"name"`
	EmpID           int    `json:"emp_id"`
	EmployeeDetails `json:"det"`
}

func (e *EmployeeDetails) String() string {
	return fmt.Sprintf("Custom String\nEmail : %10s\nAddress : %10s\nPhoneno : %10d\n", e.Email, e.Address, e.PhoneNo)
}

func main() {
	var emp1 = Employee{Name: "Logesh", EmpID: 1,
		EmployeeDetails: EmployeeDetails{Email: "logesh@my.org", Address: "some addr", PhoneNo: 94389489}}
	fmt.Println(&(emp1.EmployeeDetails))
	empD := EmployeeDetails{Email: "some@mail.com", Address: "adders", PhoneNo: 94287187}
	fmt.Println(&empD)
}

// fmt.Println("Email   :", emp1.Email)
// fmt.Println("Address :", emp1.Address)
// fmt.Println("PhoneNo :", emp1.PhoneNo)
// fmt.Println("Email   :", emp1.EmployeeDetails.Email)
// fmt.Println("Address :", emp1.EmployeeDetails.Address)
// fmt.Println("PhoneNo :", emp1.EmployeeDetails.PhoneNo)
// fmt.Printf("%#v\n", emp1)

// jsonBytes, _ := json.MarshalIndent(emp1, "", "  ")
// fmt.Println(string(jsonBytes))
