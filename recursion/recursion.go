package recursion

import "fmt"

func main() {
	fact := factorial(7)
	fmt.Println(fact)

}
func factorial(number int) int {

	// recursion method for factorial function

	if number == 0 {
		return 1
	}
	return number * factorial(number-1)

	//Normal looping method for factorial

	//result := 1
	//for i := 1; i <= number; i++ {
	//	result = result * i
	//}
	//return result
}
