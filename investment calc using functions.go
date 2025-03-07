package main

import (
	"fmt"
	"math"
)

const inflationrate = 6.5

func main() {
	var investmentAmount float64
	var expectedReturnRate = 5.5
	var year float64

	//fmt.Print("investmentAmount :")
	output("investment amout: ")
	fmt.Scan(&investmentAmount)

	//fmt.Print("expected return rate :")
	output("expected return rate : ")
	fmt.Scan(&expectedReturnRate)

	//fmt.Print("Year:")
	output("Year")
	fmt.Scan(&year)
	futurevalue, realfuturevalue := calculations(investmentAmount, expectedReturnRate, year)

	//futurevalue := investmentAmount * math.Pow(1+expectedReturnRate/100, year)
	//realfuturevalue := futurevalue / math.Pow(1+inflationrate/100, year)

	//using printf statement

	fmt.Printf("futuure value:%.2f\nReal Future Value :%.2f\n", realfuturevalue, futurevalue) // printing the 2 outputs in the single fmt.printf
	//fmt.Printf("Real Future Value :%v\n", realfuturevalue)                                 //this printing only one we can use both method to print the output

	//using println statement
	//fmt.Println("future value: ", futurevalue)
	//fmt.Println("Real Future Value is : ", realfuturevalue)clear
}
func output(text string) {
	fmt.Print(text)

}
func calculations(investmentAmount, expectedReturnRate, year float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, year)
	rfv = fv / math.Pow(1+inflationrate/100, year)
	return fv, rfv

}
