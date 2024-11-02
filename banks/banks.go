package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalancefromFile() float64 {

	balance, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 0.0
	}

	accountBalance, _ := strconv.ParseFloat(string(balance), 64)
	return accountBalance
}
func writeToFile(accountBalance float64) {
	os.WriteFile(accountBalanceFile, []byte(fmt.Sprintf("%.2f", accountBalance)), 0666)
}
func main() {
	accountBalance := getBalancefromFile()
	for {
		fmt.Println("Welcome to the bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check your balance: ")
		fmt.Println("2. Deposit: ")
		fmt.Println("3. Withdraw: ")
		fmt.Println("4. Exit: ")

		var choice int
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1
		// wantsDeposit := choice == 2
		// wantsWithdraw := choice == 3
		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Println("How much do you want to deposit?")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Your new balance is: ", accountBalance)
		case 3:
			fmt.Println("How much do you want to withdraw?")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount > accountBalance {
				fmt.Println("You don't have enough money")
				continue
			}

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Your new balance is: ", accountBalance)
		default:
			fmt.Println("Thank you for Banking with us. Goodbye!")
			return
		}
		writeToFile(accountBalance)
	}
	// if wantsCheckBalance {
	// 	fmt.Println("Your balance is: ", accountBalance)
	// } else if wantsDeposit {
	// 	fmt.Println("How much do you want to deposit?")
	// 	var depositAmount int
	// 	fmt.Scan(&depositAmount)
	// 	if depositAmount <= 0 {
	// 		fmt.Println("Invalid amount")
	// 		continue
	// 	}
	// 	accountBalance += depositAmount
	// 	fmt.Println("Your new balance is: ", accountBalance)
	// } else if wantsWithdraw {
	// 	fmt.Println("How much do you want to withdraw?")
	// 	var withdrawAmount int
	// 	fmt.Scan(&withdrawAmount)

	// 	if withdrawAmount > accountBalance {
	// 		fmt.Println("You don't have enough money")
	// 		continue
	// 	}

	// 	if withdrawAmount <= 0 {
	// 		fmt.Println("Invalid amount")
	// 		continue
	// 	}

	// 	accountBalance -= withdrawAmount
	// 	fmt.Println("Your new balance is: ", accountBalance)
	// } else {
	// 	fmt.Println("Thank you for Banking with us. Goodbye!")
	// 	return
	// }

}
