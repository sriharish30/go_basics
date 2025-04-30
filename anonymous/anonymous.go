package anonymous

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4}

	double := createTransform(2)
	triple := createTransform(3)

	// using anonymous function method

	transformednumbers := transformedNumbers(numbers, func(number int) int {
		return number * 4
	})
	doubled := transformedNumbers(numbers, double)
	tripled := transformedNumbers(numbers, triple)

	fmt.Println(transformednumbers)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformedNumbers(numbers []int, transform func(int) int) []int {

	doubleNumber := []int{}

	for _, value := range numbers {
		doubleNumber = append(doubleNumber, transform(value))

	}
	return doubleNumber

}

//Understanding clousures

func createTransform(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
