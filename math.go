package main

import (
	"fmt"
	// "math"
)

// Documentation help: https://pkg.go.dev/fmt

// Greatest Common Denominator
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// Least Common Denominator
func lcd(a, b int) int {
	return (a * b) / gcd(a, b)
}

// Basic fraction formula, something along lines of a/b +/-/* or / c/d
func fractionSimplifier() {
	var numer1, numer2, denom1, denom2 int
	var sign string

	fmt.Println("Enter numbers for the following format: A/B +/- C/D")
	fmt.Print("Enter first fraction: (ie: 25 30) ")
	fmt.Scanf("%d", &numer1)
	fmt.Scanf("%d", &denom1)
	fmt.Print("Enter second fraction: (ie: 35 12) ")
	fmt.Scanf("%d", &numer2)
	fmt.Scanf("%d", &denom2)
	fmt.Print("Enter sign: (ie: +, -, *, /) ")
	fmt.Scanf("%s", &sign)

	// search for proper contol input type validation.
	// Put in scenario for when the value is negative
	switch sign {
	case "+":
	case "-":
	case "*":
	case "/":
	default:
		fmt.Println("Invalid sign, please try again.")
		fractionSimplifier()
	} 
}


// Pre Alegbra/ Elementary Algebra functions
// Will use %g for floats
// func average() {

// }

// func probability() {

// }

// func quadraticFormula() {

// }


/* Main will be used to call each section as necessary ranging from:
1. Pre-Algebra / Elementary Algebra
2. Intermediate Algebra / Coordinate Geometry 
3. Plane Geometry
4. Trigonometry */
func main() {
	fractionSimplifier()

}