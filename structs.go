package main

import (
	"fmt"
	"time"
)

type user struct {
	firstname string
	lastname  string
	birth     string
	age       string
	createdOn time.Time
}

func userinput(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}
func output(u user) {
	fmt.Println(u.firstname, u.lastname, u.birth, u.age, u.createdOn)
}
func main() {
	userfirstname := userinput("enter the first name:")
	userlastname := userinput("enter the last name:")
	userbirthdate := userinput("enter the date:")
	userage := userinput("enter the age:")

	//var userDetails user
	userDetails := user{
		firstname: userfirstname,
		lastname:  userlastname,
		birth:     userbirthdate,
		age:       userage,
		createdOn: time.Now(),
	}
	output(userDetails)

}
