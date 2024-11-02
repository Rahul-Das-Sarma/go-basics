package main

import (
	"fmt"
	"math"
)

var inflationRate = 6.5

func main() {
	var investmentAmount float64
	expectedInterestRate := 8.5
	var years float64 = 10
	outputPrint("Enter investment amount: ")
	fmt.Scan(&investmentAmount)
	// var futureValue = investmentAmount * math.Pow((1+(expectedInterestRate*0.01)), years)
	// var futureRealvalue = futureValue / math.Pow(1+inflationRate/100, years)
	var futureValue, futureRealvalue = calcFutureValue(investmentAmount, expectedInterestRate, years)

	// fmt.Println("Future value: ", futureValue, futureRealvalue)
	fv := fmt.Sprintf("Future value: %.2f", futureValue)
	rfv := fmt.Sprintf("Future real value: %.2f", futureRealvalue)
	fmt.Println(fv, rfv)
}

func outputPrint(text string) {
	fmt.Print(text)
}

func calcFutureValue(investmentAmount float64, expectedInterestRate float64, years float64) (futureValue float64, futureRealvalue float64) {
	futureValue = investmentAmount * math.Pow((1+(expectedInterestRate*0.01)), years)
	futureRealvalue = futureValue / math.Pow(1+inflationRate/100, years)

	return futureValue, futureRealvalue
}
