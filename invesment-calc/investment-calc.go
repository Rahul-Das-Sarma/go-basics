package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64
	expectedInterestRate := 8.5
	var years float64 = 10
	inflationRate := 6.5
	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)
	var futureValue = investmentAmount * math.Pow((1+(expectedInterestRate*0.01)), years)
	var futureRealvalue = futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Future value: ", futureValue, futureRealvalue)
}
