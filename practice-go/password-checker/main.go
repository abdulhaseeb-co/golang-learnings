package main

import (
	"fmt"
	"unicode"
)

// Function to check password strength
func passwordChecker(pw string) bool {
	// Convert the password string into a slice of runes (supports multi-byte characters)
	pwR := []rune(pw)

	// Check if password is at least 8 characters long
	if len(pwR) < 8 {
		return false
	}

	// Boolean flags to check for required character types
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	// Loop through each character in the password
	for _, v := range pwR {
		if unicode.IsUpper(v) {
			hasUpper = true
		}
		if unicode.IsLower(v) {
			hasLower = true
		}
		if unicode.IsNumber(v) {
			hasNumber = true
		}
		if unicode.IsPunct(v) || unicode.IsSymbol(v) {
			hasSymbol = true
		}
	}

	// Return true only if all conditions are met
	return hasUpper && hasLower && hasNumber && hasSymbol
}

func main() {
	// Test with an invalid password
	if passwordChecker("") {
		fmt.Println("password good")
	} else {
		fmt.Println("password bad")
	}

	// Test with a valid password
	if passwordChecker("This!I5A") {
		fmt.Println("password good")
	} else {
		fmt.Println("password bad")
	}

	// Test with another invalid password
	if passwordChecker("password") {
		fmt.Println("password good")
	} else {
		fmt.Println("password bad")
	}

	// Test with a strong password
	if passwordChecker("Str0ng!Pass") {
		fmt.Println("password good")
	} else {
		fmt.Println("password bad")
	}
}
