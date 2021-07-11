package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"simpleGoProgram/accounts"
)

func main() {

	defaultWithdrawLimit := accounts.WithDrawLimit{
		LastWithDrawDate: time.Now().Add(24 * time.Hour),
		LastLimitSet: 300,
		Amount: 300,
	}

	currentAccount := accounts.CurrentAccount{
		Name:          "Current",
		PAN:           "1234 5678 9101 1213",
		AccountNumber: "87598003",
		SortCode:      "75-34-09",
		Balance: 0.0,
		Statements:    make([]accounts.Statement,0,10), WithDrawLimit: defaultWithdrawLimit,
	}

	savingsAccount := accounts.SavingsAccount{
		Name:          "Current",
		PAN:           "1234 5678 9101 1213",
		AccountNumber: "87598003",
		SortCode:      "75-34-09",
		Balance: 0.0,
		Statements:    make([]accounts.Statement,0,10), WithDrawLimit: defaultWithdrawLimit,
	}

	// initialized empty map of type iAccount interface
	accountCards := make(map[string]accounts.IAccount)
	accountCards["current"] = &currentAccount
	accountCards["savings"] = &savingsAccount

	options := func() string {
		return "::::Choose an account to proceed or show accountCards::::\nEnter\n" +
			"0: current account\n" +
			"1: savings account\n" +
			"2: show accountCards\n" +
			"3: Exit program\n"
	}

	cardSelected := -1
	for {
		restart : // define label to be used with goto statement
		fmt.Println(options())
		fmt.Scan(&cardSelected)

		switch cardSelected {
		case 0: val := startApp(accountCards["current"])
			if strings.Compare(val, "restart") == 0 {
				goto restart
			}
			break
		case 1: val := startApp(accountCards["savings"])
			if strings.Compare(val, "restart") == 0 {
				goto restart
			}
			break
		case 2:listAccounts(accountCards)
			break
		case 3: os.Exit(0)

		default:
			fmt.Println("Invalid option!")
		}
	}

}

func listAccounts(account map[string]accounts.IAccount)  {
	for key, element := range account{
		fmt.Println("==============",key,"==============")
		element.PrintAccountDetails()
		fmt.Println("............................................")
	}
}

func startApp(account accounts.IAccount) string{
	option := 0
	displayInstructions()
	for option != 8 {
		fmt.Scan(&option)

		switch option {
		case 0: displayInstructions()
			break

		case 1: account.Deposit(inputAmount())
			break

		case 2: account.Withdraw(inputAmount())
			break

		case 3: account.ShowBalance()
			break

		case 4: account.SetWithdrawLimit(inputAmount())
			break

		case 5: account.ShowStatement()
			break

		case 6: account.PrintOutStatement()
			break

		case 7: return "restart"
		case 8: os.Exit(0)
		default: fmt.Println("Invalid selection!")
		}
	}
	return ""
}

func inputAmount() float64{
	fmt.Println("Type in the amount:")
	amount := 0.0
	fmt.Scan(&amount)
	return amount
}

func displayInstructions(){
	fmt.Printf("::::::::::::::::::Available actions:::::::::::::::::\nEnter\n" +
		"0: To print available options\n" +
		"1: To make a deposit\n" +
		"2: To withdraw\n" +
		"3: To show balance\n" +
		"4: To set withdrawal limit per day\n" +
		"5: To show statments\n" +
		"6: To print out statement in CVS format\n" +
		"7: To return to main menu\n" +
		"8: To exit program\n" +
		"Select an action:")
}
