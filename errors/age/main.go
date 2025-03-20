package main

import (
	"errors"
	"fmt"
)

// 1. Division with error handling
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// 2. Custom error type for age validation
type AgeError struct {
	Age     int
	Message string
}

func (e *AgeError) Error() string {
	return fmt.Sprintf("%s: %d", e.Message, e.Age)
}

func validateAge(age int) error {
	if age < 0 {
		return &AgeError{Age: age, Message: "age cannot be negative"}
	}
	if age > 150 {
		return &AgeError{Age: age, Message: "age seems too high"}
	}
	return nil
}

// 3. File operations with multiple error checks
func processUserData(name string, age int) error {
	// Check name
	if name == "" {
		return errors.New("name cannot be empty")
	}

	// Validate age
	if err := validateAge(age); err != nil {
		return fmt.Errorf("invalid age: %w", err)
	}

	return nil
}

func main() {
	// Test division
	if result, err := divide(10, 0); err != nil {
		fmt.Printf("Division error: %v\n", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}

	// Test age validation
	if err := validateAge(200); err != nil {
		if ageErr, ok := err.(*AgeError); ok {
			fmt.Printf("Age validation failed: %v\n", ageErr)
		}
	}

	// Test user data processing
	if err := processUserData("", 25); err != nil {
		fmt.Printf("Processing error: %v\n", err)
	}
}
