package main

import "fmt"

type Account struct {
	balance float64
}

func (a *Account) withdraw(amount float64) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Transaction failed: %v\n", r)
		}
	}()

	if amount <= 0 {
		panic("withdrawal amount must be positive")
	}

	if amount > a.balance {
		panic(fmt.Sprintf("insufficient funds: balance %.2f, withdrawal %.2f", a.balance, amount))
	}

	a.balance -= amount
	fmt.Printf("Withdrew %.2f. New balance: %.2f\n", amount, a.balance)
}

func main() {
	account := &Account{balance: 100}

	// Test valid withdrawal
	fmt.Println("Attempting valid withdrawal...")
	account.withdraw(50)

	// Test invalid amount
	fmt.Println("\nAttempting negative withdrawal...")
	account.withdraw(-30)

	// Test insufficient funds
	fmt.Println("\nAttempting withdrawal exceeding balance...")
	account.withdraw(1000)

	fmt.Println("\nProgram continues to run!")
}
