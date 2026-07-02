package main

import "fmt"

// What is defer?
// The defer statement postpones the execution of a function until the surrounding function returns. 
// It's used to ensure that a function call is performed later, no matter what happens in the meantime.

// Key Characteristics:
// Delayed Execution: Code runs after the current function completes
// LIFO Order (Last In, First Out): If multiple defers exist, they execute in reverse order
// Arguments Evaluated Immediately: The arguments are evaluated when defer is encountered, not when it executes
// Common Use Cases:
// Closing files or database connections
// Unlocking mutexes
// Cleanup operations
// Exception handling


// defer is used to delay the execution of a function until the surrounding 
// function returns. The deferred call's arguments are evaluated immediately, 
// but the function call is not executed until the surrounding function returns.

func main() {
	// Example 1: Simple defer
	fmt.Println("--- Example 1: Simple Defer ---")
	fmt.Println("Start")
	defer fmt.Println("Deferred: This executes last")
	fmt.Println("End")

	fmt.Println("\n--- Example 2: Multiple Defers (LIFO) ---")
	defer fmt.Println("Defer 1 (Last In)")
	defer fmt.Println("Defer 2 (Middle)")
	defer fmt.Println("Defer 3 (First In - executes first)")

	fmt.Println("Main execution")

	fmt.Println("\n--- Example 3: Defer with Arguments ---")
	value := 10
	defer printValue("Deferred value:", value)
	value = 20
	fmt.Println("Changed value to:", value)
}

func printValue(label string, val int) {
	fmt.Printf("%s %d\n", label, val)
}
