package main

import "fmt"

func main() {

	newlist := []float64{23.4, 56.7, 89.6}
	fmt.Println(newlist)
	merge := []float64{90.8, 78.78, 123.09}
	newlist = append(newlist, merge...)
	fmt.Println(newlist)
}
