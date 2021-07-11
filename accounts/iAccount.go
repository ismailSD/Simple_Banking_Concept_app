package accounts
type IAccount interface {
	Deposit(amount float64)
	Withdraw(amount float64)
	ShowBalance()
	SetWithdrawLimit(amount float64) // per day
	ShowStatement()
	PrintOutStatement()
	PrintAccountDetails()
}