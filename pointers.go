package main

import "fmt"

func main() {
	age := 32
	//var address *int

	address := &age
	fmt.Println("Age: ", *address)
	getadultage(address)
	fmt.Println(age)
}

func getadultage(age *int) {
	*age = *age - 18

}
