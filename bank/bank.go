package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/dreking/bank/utils"
)

const accountBalanceFileName = "balance.txt"

func main() {
	accountBalance, err := utils.GetFloatFromFile(accountBalanceFileName)
	if err != nil {
		fmt.Println("Error:", err)
		// return
		panic("Cant continue, sorry!")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach out to us on", randomdata.PhoneNumber())

	for {
		presentOptions()

		var choice int
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your account balance is:", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Your new account balance is:", accountBalance)
			utils.WriteFloatToFile(accountBalance, accountBalanceFileName)
		} else if choice == 3 {
			fmt.Print("Your withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient balance")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Your new account balance is:", accountBalance)
			utils.WriteFloatToFile(accountBalance, accountBalanceFileName)
		} else {
			fmt.Println("See you next time")
			// return
			break
		}
	}

	fmt.Println("Thanks for choosing our bank!")

}
