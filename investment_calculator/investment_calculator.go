package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	// Get user input from CLI
	outputText("Enter investment amount:")
	fmt.Scan(&investmentAmount)

	outputText("Enter expected return rate:")
	fmt.Scan(&expectedReturnRate)

	outputText("Enter number of years:")
	fmt.Scan(&years)

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	futureValue, futureRealValue := calculateReturnValues(investmentAmount, expectedReturnRate, years)

	// fmt.Printf("Future Value: %.1f\nFuture Value(adjisted for Inflation): %.1f\n", futureValue, futureRealValue)
	// fmt.Println("Future Value(adkisted for Inflation)", futureRealValue)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value(adjisted for Inflation): %.1f\n", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Println(text)
}

// Function that returns values without return initialization
func calculateReturnValues(investmentAmount float64, expectedReturnRate float64, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)

	return fv, rfv
}

// Function that returns values with return initialization
// func calculateReturnValues(investmentAmount float64, expectedReturnRate float64, years float64) (fv float64, rfv float64) {
// 	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	rfv = fv / math.Pow(1+inflationRate/100, years)

// 	return
// }
