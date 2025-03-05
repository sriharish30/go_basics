package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxrate float64
	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxrate)
	Ebt := revenue - expenses
	profit := Ebt * (1 - taxrate/100)
	ratio := Ebt / profit
	fmt.Println("Earning Before Tax: ", Ebt)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)

}
