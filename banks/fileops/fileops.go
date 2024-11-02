package fileops

import (
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(file string) float64 {

	balance, err := os.ReadFile(file)
	if err != nil {
		return 0.0
	}

	accountBalance, _ := strconv.ParseFloat(string(balance), 64)
	return accountBalance
}
func WriteToFile(value float64, fileName string) {
	os.WriteFile(fileName, []byte(fmt.Sprintf("%.2f", value)), 0666)
}
