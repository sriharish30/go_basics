package main

import "fmt"

func main() {
	prices := []int{2, 5, 7, 9, 6}
	prices = append(prices, 1, 8)
	prices = prices[2:]

	fmt.Println(prices)
}

//example practices
//func main() {
//	var product [4]int = [4]int{23, 45, 67, 89}
//	fmt.Println("product:", product)

//	prices := [5]int{12, 45, 89, 67}

//	fmt.Println("prices:", prices[2])
//	prices[4] = 123
//	fmt.Println("new price:", len(prices), cap(prices))
//	featuredPrice := prices[3:]
//	fmt.Println(featuredPrice)
//	fmt.Println("featuredPrice:", len(featuredPrice), cap(featuredPrice))

//}
