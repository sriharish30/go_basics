package main

import "fmt"

//func main() {

// it is a normal method for sum the values using for loop

//numbers := []int{1, 10, 15}
//sum := sumup(numbers)
//fmt.Println(sum)
//}
//func sumup(numbers []int) int {
//sum := 0
//for _, value := range numbers {
//	sum = sum + value
//}
//return sum

//}

func main() {

	// it is the variadic function for sum the values using for loop

	sum := sumup(12, 3, 5)
	fmt.Println(sum)
	numbers := []int{1, 10, 15}
	anothersum := sumup(numbers...)
	fmt.Println(anothersum)
}
func sumup(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum = sum + value
	}
	return sum

}
