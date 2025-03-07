package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxrate float64

	//fmt.Print("Revenue: ")
	output("Revenue: ")
	fmt.Scan(&revenue)
	//fmt.Print("Expenses: ")
	output("expenses: ")
	fmt.Scan(&expenses)
	//fmt.Print("Tax Rate: ")
	output("TaxRate: ") 
	fmt.Scan(&taxrate)

	Ebt, profit, ratio := profitcalculation(revenue, expenses, taxrate)
	//Ebt := revenue - expenses
	//profit := Ebt * (1 - taxrate/100)
	//ratio := Ebt / profit
	fmt.Println("Earning Before Tax :", Ebt)
	fmt.Printf("Profit:%.2f\n", profit)
	fmt.Printf("Ratio:%.2f\n", ratio)

}
func output(text string) {
	fmt.Print(text)

}
func profitcalculation(revenue, expenses, taxrate float64) (Ebt1 float64, profit1 float64, ratio1 float64) {
	Ebt1 = revenue - expenses
	profit1 = Ebt1 * (1 - taxrate/100)
	ratio1 = Ebt1 / profit1
	return Ebt1, profit1, ratio1
}
