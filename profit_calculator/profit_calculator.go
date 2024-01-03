package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserText("Enter a revenue")
	if err != nil {
		panic(err)
	}
	expenses, err := getUserText("Enter expenses")
	if err != nil {
		panic(err)
	}
	taxRate, err := getUserText("Enter a tax rate")
	if err != nil {
		panic(err)
	}

	// Earnings Before Tax
	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("ebt: %.1f\n", ebt)
	fmt.Printf("Profit %.1f\n", profit)
	fmt.Printf("Ratio: %.1f\n", ratio)

}

func getUserText(text string) (float64, error) {
	var userInput float64
	fmt.Println(text)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("number is invalid")
	}
	return userInput, nil
}

func calculateFinancials(revenue float64, expenses float64, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	storeResultsToFile(ebt, profit, ratio)
	return ebt, profit, ratio
}

func storeResultsToFile(ebt float64, profit float64, ratio float64) {
	ebtString := fmt.Sprint(ebt)
	profitString := fmt.Sprint(profit)
	ratioString := fmt.Sprint(ratio)
	data := ebtString + ":" + profitString + ":" + ratioString

	os.WriteFile("profit.txt", []byte(data), 0644)
}
