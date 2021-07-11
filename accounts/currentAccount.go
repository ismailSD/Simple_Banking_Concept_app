package accounts

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type CurrentAccount struct {
	Name string// account type e.g., debit
	PAN string//16 digit Permanent Account Number
	AccountNumber string// 8 digit account number
	SortCode string// 6 digit sort code

	Balance       float64
	Statements    []Statement
	WithDrawLimit WithDrawLimit
}

// PrintAccountDetails ============= CURRENT ACCOUNT FUNCTIONALITIES =========================
func (account *CurrentAccount) PrintAccountDetails() {
	fmt.Printf("Name: %s \nPAN: %s \nAccount Number: %s \nSort Code: %s \n", account.Name, account.PAN, account.AccountNumber,
		account.SortCode)
}

func (account *CurrentAccount) Deposit(amount float64) {
	// 1: ensure that Amount is not less than zero or greater than one million
	// 2: add the Amount to existing Balance
	// 3: create a statement that includes the current Date and time
	// 4: ensure that Amount is not less than zero or greater than one million

	if amount <=0 || amount > 1000000 {
		fmt.Println("Amount must be greater 0 and less than or equal to 1,000,000")
		return
	}

	account.Balance += amount
	fmt.Println("Deposit of (",amount,") success!")

	// create a statement after each Deposit
	account.Statements = append(account.Statements, Statement{time.Now().Format("2006-01-02 3:4:5 PM"),
		"Deposit", amount, 0.0, account.Balance})
}

func (account *CurrentAccount) Withdraw(amount float64) {
	// 1: ensure that there is sufficient funds
	// 2: ensure that Amount is not less than zero or greater than one million
	// 3: ensure that Amount is not greater than the account Balance
	// 4: check if 24 hours has past since last with draw limit set if so renew it
	// 5: check if Withdraw limit has reached
	// 6: ensure that Amount is not greater than Withdraw limit Amount
	// 7: subtract Amount from account Balance
	// 8: Create a statement that includes the current Date and time

	if account.Balance <=0 {
		fmt.Println("Insufficient funds!")

	} else if amount <=0 || amount > 1000000{
		fmt.Println("Amount must be greater 0 and less than or equal to 1,000,000")

	} else if amount > account.Balance {
		fmt.Println("Cannot Withdraw more than available Balance of (",account.Balance,")!")

	} else if account.WithDrawLimit.LastWithDrawDate.Before(time.Now()) {
		// 24 hours has past reset the Withdraw limit per day
		// rest the Withdraw Amount per day
		account.WithDrawLimit.LastWithDrawDate = time.Now().Add(24 * time.Hour)
		account.WithDrawLimit.Amount = account.WithDrawLimit.LastLimitSet

	} else if account.WithDrawLimit.Amount <= 0{
		fmt.Println("You have reached the Withdraw limit for today!")

	} else if amount > account.WithDrawLimit.Amount {
		fmt.Println("You are trying to Withdraw over the limit per day!")
		fmt.Println("You can with draw ", account.WithDrawLimit.Amount, " or less.")

	}else {
		account.Balance -= amount
		account.WithDrawLimit.Amount -= amount
		fmt.Println("Withdraw of (",amount,") success!")

		// create a statement after each Withdraw
		account.Statements = append(account.Statements,
			Statement{time.Now().Format("2006-01-02 3:4:5 PM"), "Withdraw",
				0.0, amount, account.Balance})
	}
}

func (account *CurrentAccount) ShowBalance()  {
	fmt.Printf("Available Balance: %.2f \n",account.Balance)
}

func (account *CurrentAccount) SetWithdrawLimit(amount float64) {
	// 1: ensure that the limit Amount is not less than 1 or greater one thousand
	// 2: assign account limit to the new Amount

	if amount < 1 || amount > 1000{
		fmt.Println("Limit must be between 1 and 1,000")

	}else {
		account.WithDrawLimit.Amount = amount
		account.WithDrawLimit.LastWithDrawDate = time.Now().Add(24 * time.Hour)
		fmt.Println("Withdraw limit set to", account.WithDrawLimit.Amount, "per day.")
	}

}

func (account *CurrentAccount) ShowStatement() {
	// iterate through account statements and display them
	if len(account.Statements) <= 0{
		fmt.Println("There are no statements available.")
		return
	}
	fmt.Println("Date\t\tDescription\t\tDeposit\t\tWithdrawal\t\tBalance")
	fmt.Println(strings.Repeat("=",110))
	for _, s := range account.Statements {
		fmt.Printf("%s \t\t %s \t\t %.2f \t\t %.2f \t\t %.2f\n", s.Date,s.Description,s.Deposit,s.Withdrawal,s.Balance)
		fmt.Println(strings.Repeat(".",110))
	}
	fmt.Print("\n\n")
}
func (account *CurrentAccount) PrintOutStatement() {
	// 1: create csv files named statement.csv
	// 2: create statement Description header
	// 3: iterate through account statements and assign to 2D array
	// 4: iterate through 2D array and write each row data to the cvs file

	csvFile, err := os.Create("../files/current_statement.csv")
	if err != nil { log.Fatalf("Failed creating files: %s", err)}
	csvWriter := csv.NewWriter(csvFile)

	rows := make([][]string, 0,10)
	rows = append(rows, []string{ "Date", "Description", "Deposit", "Withdrawal", "Balance"})

	for _, s := range account.Statements {
		rows = append(rows, []string{
			s.Date,
			s.Description,
			fmt.Sprint(s.Deposit),
			fmt.Sprint(s.Withdrawal),
			fmt.Sprint(s.Balance),
		})
	}
	for _, row := range rows {
		_ = csvWriter.Write(row)
	}
	csvWriter.Flush()
	csvFile.Close()
	fmt.Println("Statement saved to file: ", csvFile.Name())
}