package main

import "fmt"

// Package-level constant accessible anywhere in this package
const p = 100

// The init function runs automatically before the main function execution begins
func init() {
	fmt.Println("System Initialized... Bank")
}

// 'outer' function takes an initial amount (money) and another variable (e.g., bonus).
// It returns an anonymous function of type `func(int) int` which acts as our closure.
func outer(money int, bonus int) func(int) int {

	// This anonymous function is the closure.
	// It captures 'money' and 'bonus' from the outer function's scope.
	return func(expense int) int {
		// We manipulate the captured 'money' variable.
		// The state is remembered across multiple calls.
		money = (money + bonus) - expense

		// It also has access to package-level constants like 'p'
		fmt.Printf("Transaction processed. Base constraint: %d | ", p)

		return money
	}
}

func main() {
	// We call 'outer' and it returns the inner anonymous function (the closure).
	// The 'show' variable now holds this closure function.
	// The state of 'money' (initialized to 1000) and 'bonus' (200) is preserved.
	show := outer(1000, 200)

	// Calling the closure for the first time
	// Calculation: 1000 (money) + 200 (bonus) - 300 (expense) = 900
	currentBalance := show(300)
	fmt.Printf("Current Balance: %d\n", currentBalance)

	// Calling the exact same closure a second time.
	// Notice how it REMEMBERS the 'money' from the previous execution (900).
	// Calculation: 900 (money) + 200 (bonus) - 100 (expense) = 1000
	currentBalance = show(100)
	fmt.Printf("Current Balance: %d\n", currentBalance)
}
