package main

import (
	"errors"
	"fmt"
)

type Account struct {
	AccountNumber string
	Balance       float64
	OwnerName     string
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("Deposit amount must be positive")
	}
	a.Balance += amount
	fmt.Printf("Deposited $%.2f to %s. New Balance: $%.2f\n", amount, a.AccountNumber, a.Balance)
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("Withdrawal amount must be positive")
	}
	if amount > a.Balance {
		return errors.New("Insufficient funds")
	}
	a.Balance -= amount
	fmt.Printf("Withdrew $%.2f from %s. New Balance: $%.2f\n", amount, a.AccountNumber, a.Balance)
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.Balance
}
func (a *Account) String() string {
	return fmt.Sprintf("Account: [%s] Owner: %s, Balance: $%.2f", a.AccountNumber, a.OwnerName, a.Balance)
}

type SavingsAccount struct {
	Account
	InterestRate float64
}

func (sa *SavingsAccount) AddInterest() {
	interest := sa.Balance * sa.InterestRate
	fmt.Printf("Adding interest $%.2f to savings account %s. ", interest, sa.AccountNumber)
	err := sa.Deposit(interest)
	if err != nil {
		fmt.Printf("Error depositing $%.2f to savings account. %v\n", interest, err)
	}
}

type OverdraftAccount struct {
	Account
	OverdraftLimit float64
}

func (oa *OverdraftAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("Withdrawal amount must be positive")
	}
	if (oa.Balance + oa.OverdraftLimit) < amount {
		return fmt.Errorf("withdrawal of $%.2f exceeds overdraft limit for %s,Available include overdraft:$%.2f", amount, oa.OwnerName, oa.Balance+oa.OverdraftLimit)
	}
	oa.Balance -= amount
	fmt.Printf("widrew $%.2f from overdraft account %s. New Balance: $%.2f\n", amount, oa.AccountNumber, oa.Balance)
	return nil
}

func main() {
	fmt.Println("--- Bank Statement ---")
	savAcc := SavingsAccount{
		Account: Account{
			AccountNumber: "SAV001",
			Balance:       1000.00,
			OwnerName:     "Alice Saver",
		},
		InterestRate: 0.02,
	}
	fmt.Println("\n--- Savings Account Operations ---")
	fmt.Println(savAcc.Account.String())

	err := savAcc.Deposit(200)
	if err != nil {
		fmt.Printf("Error depositing $%.2f to savings account. %v \n", 200.00, err)
	}
	fmt.Println(savAcc.Account.String())

	savAcc.AddInterest()
	fmt.Println(savAcc.Account.String())

	err = savAcc.Withdraw(50.00)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Final Savings Details:", savAcc.Account.String())
	over := OverdraftAccount{
		Account: Account{
			AccountNumber: "OVR001",
			Balance:       1000.00,
			OwnerName:     "Bob",
		},
		OverdraftLimit: 200,
	}

	fmt.Println("---Overdraft account")
	fmt.Println(over.Account.String())
	err = over.Deposit(50.00)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(over.Account.String())
	err = over.Withdraw(1300)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(over.Account.String())

}
