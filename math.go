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

	switch sign {
	case "+":
		// Find active LCD
		active_lcd := lcd(denom1, denom2)
		
		// Find the adjusted numerators that work along with the LCD via multiplication
		frac1 := numer1 * (active_lcd / denom1)
		frac2 := numer2 * (active_lcd / denom2)
		frac_total := frac1 + frac2 

		// Progression Check
		fmt.Println("Current Fraction after addition: ", frac_total, "/", active_lcd)
		
		reduced_frac_by := gcd(frac_total, active_lcd)
		reduced_numerator := frac_total / reduced_frac_by
		reduced_denominator := active_lcd / reduced_frac_by 

		// Convert to mixed numbers
		mixed_whole := reduced_numerator / reduced_denominator
		mixed_numerator := reduced_numerator % reduced_denominator

		// Solution if there is a whole number
		if mixed_whole != 0 {
			fmt.Println("Solution: ", mixed_whole, " ", mixed_numerator, "/", reduced_denominator)
		} else { // May be invalid with current output
			fmt.Println("Solution: ", reduced_numerator, "/", reduced_denominator)
		}
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