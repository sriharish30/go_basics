package person

import (
	"errors"
	"fmt"
)

type Person struct {
	name      string
	birthdate string
	jobRoll   string
}
type admin struct {
	email    string
	password string
	Person
}

func (p Person) Output() {
	fmt.Println(p.name, p.birthdate, p.jobRoll)
}
func (p *Person) Clearuser() {
	p.name = ""
	p.birthdate = ""
}
func Newadmin(email, password string) admin {
	return admin{
		email:    email,
		password: password,
		Person: Person{
			name:      "sriharish",
			birthdate: "------",
			jobRoll:   "developer",
		},
	}
}

// passing error message
func Newperson(name, birthdate, JobRoll string) (*Person, error) {
	if name == "" || birthdate == "" || JobRoll == "" {
		return nil, errors.New("name, birthdate,jobroll is required")
	}
	return &Person{
		name:      name,
		birthdate: birthdate,
		jobRoll:   JobRoll,
	}, nil

}
