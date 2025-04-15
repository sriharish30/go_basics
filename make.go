package main

import "fmt"

func main() {
	userName := make([]string, 2, 5)
	userName[0] = "sri"
	userName[1] = "s"

	userName = append(userName, "hari")
	userName = append(userName, "harishhh")

	fmt.Println(userName)

	courseratings := make(map[string]float64, 3)
	courseratings["go"] = 7.8
	courseratings["react"] = 8.91
	courseratings["angular"] = 10.
	courseratings["express"] = 90.89
	fmt.Println(courseratings)
	for index, value := range userName {
		fmt.Println("Index: ", index)
		fmt.Println("Value: ", value)
	}
	for key, value := range courseratings {
		fmt.Println("key: ", key)
		fmt.Println("Value: ", value)
	}

}
