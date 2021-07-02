package main

import (
	"fmt"
	"os"
	"strings"
)

type iAccount interface {
	deposit(amount float64)
	withdraw(amount float64)
	showBalance()
	setWithdrawLimit(amount float64)// per day
	showStatement()
	printOutStatement()
}

type statement struct {
	date string
	description string
	deposit float64
	withdrawal float64
	balance float64
}

type savingsAccount struct {
	name string// account type e.g., savings
	pAN string//16 digit Permanent Account Number
	accountNumber string// 8 digit account number
	sortCode string// 6 digit sort code

	balance float64
	statements []statement
	withDrawLimit float64
}

type currentAccount struct {
	name string// account type e.g., debit
	pAN string//16 digit Permanent Account Number
	accountNumber string// 8 digit account number
	sortCode string// 6 digit sort code

	balance float64
	statements []statement
	withDrawLimit float64
}

func (account *currentAccount) deposit(amount float64) {
	if amount <=0 {
		fmt.Println("Amount must be greater than 0")
		return
	}
	account.balance += amount
	fmt.Println("Deposit of (",amount,") success!")

	// create a statement after each deposit
	account.statements = append(account.statements,
		statement{"2020", "Deposit", amount, 0.0, account.balance})
}
func (account *currentAccount) withdraw(amount float64) {
	if amount <=0 {
		fmt.Println("Amount must be greater than 0")
		return
	}else if amount > account.balance {
		fmt.Println("Cannot withdraw more than available balance of (",account.balance,")!")
		return
	}
	account.balance -= amount
	fmt.Println("Withdraw of (",amount,") success!")

	// create a statement after each withdraw
	account.statements = append(account.statements,
		statement{"2020", "Withdraw", amount, 0.0, account.balance})
}
func (account *currentAccount) showBalance()  {
	fmt.Printf("Available balance: %.2f \n",account.balance)
}
func (account *currentAccount) setWithdrawLimit(amount float64) {
	account.withDrawLimit = amount
}
func (account *currentAccount) showStatement() {
	if len(account.statements) <= 0{
		fmt.Println("There are no statements available.")
		return
	}
	fmt.Println("Date\t\tDescription\t\tDeposit\t\tWithdrawal\t\tBalance")
	fmt.Println(strings.Repeat("=",110))
	for _, s := range account.statements {
		fmt.Printf("%s \t\t %s \t\t %.2f \t\t %.2f \t\t %.2f\n",s.date,s.description,s.deposit,s.withdrawal,s.balance)
		fmt.Println(strings.Repeat(".",110))
	}
	fmt.Print("\n\n")
}
func (account *currentAccount) printOutStatement() {
	fmt.Println("statements coming soon....")
}

func main() {

	statements := make([]statement, 0, 10)
	
	newStatement := statement{"2020", "description", 0.0, 0.0, 0.0}
	statements = append(statements, newStatement)

	account := currentAccount{
		"Current",
		"1234 5678 9101 1213",
		"87598003",
		"75-34-09",
		0.0,
		statements,
		0.0,
	}
	fmt.Println(account)
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++")
	account.deposit(100)
	account.withdraw(73)
	account.showBalance()
	account.setWithdrawLimit(100)
	account.showStatement()
	account.printOutStatement()
	option := 0
	displayInstructions()

	for option != 6{
		fmt.Scan(&option)

		switch option {
		case 0: displayInstructions()
			break
		case 1:
			account.deposit(inputAmount())
			break
		case 2:
			account.withdraw(inputAmount())
			break
		case 3: account.showBalance()
			break
		case 4: account.setWithdrawLimit(20)
			break
		case 5: account.showStatement()
			break
		case 6:account.printOutStatement()
			break
		case 7: os.Exit(0)
		default:
			fmt.Println("Invalid selection!")
		}
	}


}
func inputAmount() float64{
	fmt.Println("Type in the amount:")
	amount := 0.0
	fmt.Scan(&amount)
	return amount
}
func displayInstructions(){
	fmt.Printf("Press\n" +
		"0: To print available options\n" +
		"1: To make a deposit\n" +
		"2: To withdraw\n" +
		"3: To show balance\n" +
		"4: To set withdrawal limit per day\n" +
		"5: To show statments\n" +
		"6: To print out statement in CVS format\n" +
		"7: exit program\n" +
		"Select an action:")
}
