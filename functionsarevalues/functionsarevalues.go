package functionsarevalues

import (
	"fmt"
)

//custom type we can use this also If u can use this give type place in transformfn like this (transform is the parameter and  transformfn is the type)
//type transformfn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumber := []int{5, 1, 3}
	doubled := transformedNumbers(&numbers, double)
	tripled := transformedNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transformedfn1 := getTransformfn(numbers)
	transformedfn2 := getTransformfn(moreNumber)

	transformednumbers := transformedNumbers(&numbers, transformedfn1)
	moretransformednumbers := transformedNumbers(&moreNumber, transformedfn2)

	fmt.Println(transformednumbers)
	fmt.Println(moretransformednumbers)
}

//use custom type or otherwise use this method (transform func(int) int)

func transformedNumbers(numbers *[]int, transform func(int) int) []int {

	doubleNumber := []int{}
	for _, value := range *numbers {
		doubleNumber = append(doubleNumber, transform(value))

	}
	return doubleNumber

}
func double(number int) int {
	return number * 2
}
func triple(number int) int {
	return number * 3
}
func getTransformfn(number []int) func(int) int {
	if number[0] == 1 {
		return double

	} else {
		return triple
	}

}
