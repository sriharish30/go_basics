package main

import (
	"fmt"
)

// creating a struct

type person struct {
	name        string
	birthdate   string
	phoneNumber int
	age         int
	JobRoll     string
}

//getting the integer value in the function

func getIntData(prompt string) int {
	fmt.Print(prompt)
	var input int
	fmt.Scanln(&input)
	return input
}

//getting the string value in the function

func getstringData(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}

//creating the method for the struct

func (p person) output() {
	fmt.Println(p.name, p.birthdate, p.phoneNumber, p.age, p.JobRoll)
}

//creating the method for the struct to clear the name and birthdate value  and store it into the  main struct

func (p *person) clearuser() {
	p.name = ""
	p.birthdate = ""
}

//creating the method of the struct to create the new person and it is the pattern for creating the user

func newPerson(name string, birthdate string, phoneNumber int, age int, JobRoll string) *person {
	return &person{
		name:        name,
		birthdate:   birthdate,
		phoneNumber: phoneNumber,
		age:         age,
		JobRoll:     JobRoll,
	}

}

// main function
func main() {
	username := getstringData("Enter The Name:")
	userbirthdate := getstringData("Enter The BirthDate:")
	userphoneNumber := getIntData("Enter The Mobile Number:")
	userage := getIntData("Enter The Age:")
	userjobrole := getstringData("Enter The JobRoll:")
	person1 := newPerson(username, userbirthdate, userphoneNumber, userage, userjobrole)

	person1.output()    // it will give the output of the funtion
	person1.clearuser() //it willl clear the name and birthdate of the value given by the user
	person1.output()    //the it will print the output without the values of name and birthdate and they can print other values of job ,age like that.

}
