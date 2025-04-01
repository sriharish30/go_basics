package main

import (
	"example1/struct/person"
	"fmt"
)

func main() {
	username := GetstringData("Enter The Name:")
	userbirthdate := GetstringData("Enter The BirthDate:")
	userjobrole := GetstringData("Enter The JobRoll:")
	var person1 *person.Person
	person1, err := person.Newperson(username, userbirthdate, userjobrole)
	if err != nil {

		fmt.Println("Error:", err)
		return

	}

	admin1 := person.Newadmin("admin", "90098")
	admin1.Output()
	admin1.Clearuser()
	admin1.Output()

	person1.Output()
	person1.Clearuser()
	person1.Output()

}
