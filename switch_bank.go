package main

import "fmt"

func main() {
	var accountbalance int = 2000
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

			var depositamount int
			fmt.Print("Enter Your Deposit Amount: ")
			fmt.Scan(&depositamount)
			if depositamount <= 0 {
				fmt.Println("Invalid Amount! Deposit amount must be greater than zero")
				//return
				continue
			}

			accountbalance = depositamount + accountbalance
			fmt.Println("Your Updated AccountBalance is:", accountbalance)
		case 3:
			var withdraw int
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

		default:

			fmt.Println("Exit")
			fmt.Println("Thanks For Visiting Us")
			return

		}

	}
}
