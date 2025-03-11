package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func readfile() (float64, error) {
	data, err := os.ReadFile("balance.txt")
	if err != nil {
		return 1000, errors.New("no read file in the program")
	}
	balance := string(data)
	bal, err := strconv.ParseFloat(balance, 64)
	if err != nil {
		return 1000, errors.New("failed to parse the value")
	}

	return bal, nil
}

func writefile(balance float64) {
	bal := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(bal), 0644)

}

func main() {
	var accountbalance, err = readfile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------------------")
		//panic("sorry u r in danger now your code is crashed ")

	}
	for {
		fmt.Println("1.Account Balance: ")
		fmt.Println("2.Deposit Amount: ")
		fmt.Println("3.WithDraw Amount: ")
		fmt.Println("4.Exit")
		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Your Account Balance is: ", accountbalance)

		case 2:

			var depositamount float64
			fmt.Print("Enter Your Deposit Amount: ")
			fmt.Scan(&depositamount)
			if depositamount <= 0 {
				fmt.Println("Invalid Amount! Deposit amount must be greater than zero")
				//return
				continue
			}

			accountbalance = depositamount + accountbalance
			fmt.Println("Your Updated AccountBalance is:", accountbalance)
			writefile(accountbalance)
		case 3:
			var withdraw float64
			fmt.Print("Enter Your WithDraw Amount: ")
			fmt.Scan(&withdraw)
			if withdraw <= 0 {
				fmt.Println("Invalid Amount! The withdraw amount must be greater than 0")
				continue

			}
			if withdraw > accountbalance {
				fmt.Println("Invalid Amount ! You account dont have sufficient balance")
				continue
			}
			accountbalance = accountbalance - withdraw
			fmt.Println("Your Updated AccountBalance is:", accountbalance)
			writefile(accountbalance)

		default:

			fmt.Println("Exit")
			fmt.Println("Thanks For Visiting Us")
			return

		}

	}
}
