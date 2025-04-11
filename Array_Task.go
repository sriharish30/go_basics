package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	//task1
	hobby := [3]string{"eating", "game", "music"}

	fmt.Println(hobby)
	//task2

	fmt.Println(hobby[0])
	fmt.Println(hobby[1:3])
	//task3

	newhobbies := hobby[0:2]
	fmt.Println(newhobbies)
	//task4

	fmt.Println(cap(newhobbies))
	newhobbies = newhobbies[1:3]
	fmt.Println(newhobbies)
	//task5

	coursegoals := []string{"Learn Go", "Learn Go Basics"}
	fmt.Println(coursegoals)
	//task6

	coursegoals[1] = "Learn Go In Depth"
	coursegoals = append(coursegoals, "Learn The Go Basics")
	fmt.Println(coursegoals)
	//task7

	products := []Product{
		{
			"first-product",
			"A First Product",
			12.99,
		},
		{
			"second-product",
			"A Second Product",
			78.99,
		},
	}
	fmt.Println(products)

	newproduct := Product{
		"Third-product",
		"A Third product",
		112.45,
	}
	products = append(products, newproduct)
	fmt.Println(products)

}
