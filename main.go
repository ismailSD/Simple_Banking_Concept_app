package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
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
type withDrawLimit struct {
	lastWithDrawDate time.Time
	lastLimitSet float64
	amount float64// by default set to -1.0 means there is no limit per day
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
	withDrawLimit withDrawLimit
}

func (account *currentAccount) deposit(amount float64) {
	// 1: ensure that amount is not less than zero or greater than one million
	// 2: add the amount to existing balance
	// 3: create a statement that includes the current date and time
	// 4: ensure that amount is not less than zero or greater than one million

	if amount <=0 || amount > 1000000 {
		fmt.Println("Amount must be greater 0 and less than or equal to 1,000,000")
		return
	}

	account.balance += amount
	fmt.Println("Deposit of (",amount,") success!")

	// create a statement after each deposit
	account.statements = append(account.statements, statement{time.Now().Format("2006-01-02 3:4:5 PM"),
			"Deposit", amount, 0.0, account.balance})
}

func (account *currentAccount) withdraw(amount float64) {
	// 1: ensure that there is sufficient funds
	// 2: ensure that amount is not less than zero or greater than one million
	// 3: ensure that amount is not greater than the account balance
	// 4: check if 24 hours has past since last with draw limit set if so renew it
	// 5: check if withdraw limit has reached
	// 6: ensure that amount is not greater than withdraw limit amount
	// 7: subtract amount from account balance
	// 8: Create a statement that includes the current date and time

	if account.balance <=0 {
		fmt.Println("Insufficient funds!")

	} else if amount <=0 || amount > 1000000{
		fmt.Println("Amount must be greater 0 and less than or equal to 1,000,000")

	} else if amount > account.balance {
		fmt.Println("Cannot withdraw more than available balance of (",account.balance,")!")

	} else if account.withDrawLimit.lastWithDrawDate.Before(time.Now()) {
		// 24 hours has past reset the withdraw limit per day
		// rest the withdraw amount per day
		account.withDrawLimit.lastWithDrawDate = time.Now().Add(24 * time.Hour)
		account.withDrawLimit.amount = account.withDrawLimit.lastLimitSet

	} else if account.withDrawLimit.amount <= 0{
		fmt.Println("You have reached the withdraw limit for today!")

	} else if amount > account.withDrawLimit.amount{
		fmt.Println("You are trying to withdraw over the limit per day!")
		fmt.Println("You can with draw ", account.withDrawLimit.amount, " or less.")

	}else {
		account.balance -= amount
		account.withDrawLimit.amount -= amount
		fmt.Println("Withdraw of (",amount,") success!")

		// create a statement after each withdraw
		account.statements = append(account.statements,
			statement{time.Now().Format("2006-01-02 3:4:5 PM"), "Withdraw",
				0.0, amount, account.balance})
	}
}

func (account *currentAccount) showBalance()  {
	fmt.Printf("Available balance: %.2f \n",account.balance)
}

func (account *currentAccount) setWithdrawLimit(amount float64) {
	// 1: ensure that the limit amount is not less than 1 or greater one thousand
	// 2: assign account limit to the new amount

	if amount < 1 || amount > 1000{
		fmt.Println("Limit must be between 1 and 1,000")

	}else {
		account.withDrawLimit.amount = amount
		account.withDrawLimit.lastWithDrawDate = time.Now().Add(24 * time.Hour)
		fmt.Println("Withdraw limit set to", account.withDrawLimit.amount, "per day.")
	}

}

func (account *currentAccount) showStatement() {
	// iterate through account statements and display them
	if len(account.statements) <= 0{
		fmt.Println("There are no statements available.")
		return
	}
	fmt.Println("Date\t\tDescription\t\tDeposit\t\tWithdrawal\t\tBalance")
	fmt.Println(strings.Repeat("=",110))
	for _, s := range account.statements {
		fmt.Printf("%s \t\t %s \t\t %.2f \t\t %.2f \t\t %.2f\n", s.date,s.description,s.deposit,s.withdrawal,s.balance)
		fmt.Println(strings.Repeat(".",110))
	}
	fmt.Print("\n\n")
}
func (account *currentAccount) printOutStatement() {
	// 1: create csv files named statement.csv
	// 2: create statement description header
	// 3: iterate through account statements and assign to 2D array
	// 4: iterate through 2D array and write each row data to the cvs file

	csvFile, err := os.Create("../files/statement.csv")
	if err != nil { log.Fatalf("Failed creating files: %s", err)}
	csvWriter := csv.NewWriter(csvFile)

	rows := make([][]string, 0,10)
	rows = append(rows, []string{ "Date", "Description", "Deposit", "Withdrawal", "Balance"})

	for _, s := range account.statements {
		rows = append(rows, []string{
			s.date,
			s.description,
			fmt.Sprint(s.deposit),
			fmt.Sprint(s.withdrawal),
			fmt.Sprint(s.balance),
		})
	}
	for _, row := range rows {
		_ = csvWriter.Write(row)
	}
	csvWriter.Flush()
	csvFile.Close()
	fmt.Println("Statement saved to file: ", csvFile.Name())
}

func main() {
	account := currentAccount{
		"Current",
		"1234 5678 9101 1213",
		"87598003",
		"75-34-09",
		0.0,
		make([]statement,0,10),
		withDrawLimit{time.Now().Add(24 * time.Hour), 300, 300},
	}

	option := 0
	displayInstructions()
	for option != 7{
		fmt.Scan(&option)

		switch option {
		case 0:displayInstructions()
			break
		case 1:
			account.deposit(inputAmount())
			break
		case 2:
			account.withdraw(inputAmount())
			break
		case 3: account.showBalance()
			break
		case 4: account.setWithdrawLimit(inputAmount())
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
