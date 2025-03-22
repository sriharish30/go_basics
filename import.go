package main

import (
	"errors"
	"fmt"

	"example.com/go_practice/file"
)

func main() {

	revenue, err := getuserinput("Revenue :")
	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := getuserinput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxrate, err := getuserinput("TaxRate :")
	if err != nil {
		fmt.Println(err)
		return
	}

	Ebt, profit, ratio := mains(revenue, expenses, taxrate)

	fmt.Println("Earning Before Tax :", Ebt)
	fmt.Printf("Profit:%.2f\n", profit)
	fmt.Printf("Ratio:%.2f\n", ratio)
	file.Writefiles(Ebt, profit, ratio)

}

func getuserinput(text string) (float64, error) {
	var getinput float64
	fmt.Print(text)
	fmt.Scan(&getinput)
	if getinput <= 0 {
		return 0, errors.New("enter the  positive number")

	}
	return getinput, nil

}
