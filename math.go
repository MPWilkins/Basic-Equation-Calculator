package main

import (
	"fmt"
	// "math"
)

// Documentation help: https://pkg.go.dev/fmt

// Absolute Value for ints
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

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

// Basic fraction formula and simplification
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

	// Check which sign was entered for method of simplification
	switch sign {
	case "+":
		// Find the LCD of the denominators
		active_lcd := lcd(denom1, denom2)
		// Display the LCD being used.
		// fmt.Println("The Lowest Common Denominator is: ", active_lcd)
		
		// Adjust numberators based on LCD
		frac1 := numer1 * (active_lcd / denom1)
		frac2 := numer2 * (active_lcd / denom2)
		frac_total := frac1 + frac2 

		// Progression Check
		// fmt.Println("Fraction after addition: ", frac_total, "/", active_lcd)
		
		// Fine the GCD of the fraction
		reduced_frac_by := gcd(frac_total, active_lcd)

		// Reduce the numerator and denominator by the GCD
		reduced_numerator := frac_total / reduced_frac_by
		reduced_denominator := active_lcd / reduced_frac_by 

		// Convert to mixed numbers
		mixed_whole := reduced_numerator / reduced_denominator
		mixed_numerator := reduced_numerator % reduced_denominator

		// Solution if there is a whole number/ if the solution is a negative
		if mixed_whole != 0 && mixed_whole > 0 {
			fmt.Println("Solution: ", mixed_whole, " ", mixed_numerator, "/", reduced_denominator)
		} else if mixed_whole == 0 && reduced_numerator > 0 {
			fmt.Println("Solution: ", reduced_numerator, "/", reduced_denominator)
		} else if mixed_whole != 0 && mixed_whole < 0 {
			numerator_abs := abs(mixed_numerator)
			denominator_abs := abs(reduced_denominator)

			fmt.Println("Solution: ", mixed_whole, " ", numerator_abs, "/", denominator_abs)
		} else {
			denominator_abs := abs(reduced_denominator)

			fmt.Println("Solution: ", reduced_numerator, "/", denominator_abs)
		}
	case "-":
		// Find the LCD of the denominators
		active_lcd := lcd(denom1, denom2)
		// Display the LCD being used.
		fmt.Println("The Lowest Common Denominator is: ", active_lcd)
		
		// Adjust numberators based on LCD
		frac1 := numer1 * (active_lcd / denom1)
		frac2 := numer2 * (active_lcd / denom2)
		// Subtract the numerators
		frac_total := frac1 - frac2 

		// Progression Check
		fmt.Println("Fraction after subtraction: ", frac_total, "/", active_lcd)
		
		// Fine the GCD of the fraction
		reduced_frac_by := gcd(frac_total, active_lcd)

		// Reduce the numerator and denominator by the GCD
		reduced_numerator := frac_total / reduced_frac_by
		reduced_denominator := active_lcd / reduced_frac_by 

		// Convert to mixed numbers
		mixed_whole := reduced_numerator / reduced_denominator
		mixed_numerator := reduced_numerator % reduced_denominator

		// Solution if there is a whole number/ if the solution is a negative
		if mixed_whole != 0 && mixed_whole > 0 {
			//fmt.Println("Soulution 1")
			fmt.Println("Solution: ", mixed_whole, " ", mixed_numerator, "/", reduced_denominator)
		} else if mixed_whole == 0 && mixed_whole > 0 { // May be invalid with current output
			//fmt.Println("Soulution 2")
			fmt.Println("Solution: ", reduced_numerator, "/", reduced_denominator)
		} else if mixed_whole != 0 && mixed_whole < 0 {
			//fmt.Println("Soulution 3")
			fmt.Println("Solution: ", mixed_whole, " ", mixed_numerator * 1, "/", reduced_denominator)
		} else {
			//fmt.Println("Soulution 4") Duplicate of Solution 2?
			fmt.Println("Solution: ", reduced_numerator, "/", reduced_denominator)
		}
	case "*":
		// Find the total numerator and denominator
		new_numerator := numer1 * numer2
		new_denominator := denom1 * denom2

		active_gcd := gcd(new_numerator, new_denominator)
		// Progression Check
		fmt.Println("The GCD is:", active_gcd)

		// Reduce the numerator and denominator by the GCD
		final_numerator := new_numerator / active_gcd
		final_denominator := new_denominator / active_gcd

		// Convert to mixed numbers
		mixed_whole := final_numerator / final_denominator
		mixed_numerator := final_numerator % final_denominator

		// Display fraction
		if mixed_whole != 0 && mixed_whole > 0 {
			fmt.Println("Solution: ", mixed_whole, " ", mixed_numerator, "/", final_denominator)
		} else if mixed_whole != 0 && mixed_whole < 0 {
			numerator_abs := abs(mixed_numerator)
			denominator_abs := abs(final_denominator)
			
			fmt.Println("Solution: ", mixed_whole, " ", numerator_abs, "/", denominator_abs)
		} else if mixed_whole == 0 && (final_denominator < 0 || final_numerator < 0) { 
			numerator_abs := abs(mixed_numerator)
			denominator_abs := abs(final_denominator)

			fmt.Println("Solution: ", numerator_abs * -1, "/", denominator_abs)
		} else {
			fmt.Println("Solution: ", final_numerator, "/", final_denominator)
		}
		
	case "/":
		// Flip second fraction then multiply numerators and denominators
		new_numerator := numer1 * denom2
		new_denominator := denom1 * numer2

		active_gcd := gcd(new_numerator, new_denominator)
		// Progression Check
		fmt.Println("The GCD is:", active_gcd)

		// Reduce the numerator and denominator by the GCD
		final_numerator := new_numerator / active_gcd
		final_denominator := new_denominator / active_gcd

		// Convert to mixed numbers
		mixed_whole := final_numerator / final_denominator
		mixed_numerator := final_numerator % final_denominator

		if mixed_whole != 0 && mixed_whole > 0 {
			fmt.Println("Solution: ", mixed_whole, " ", mixed_numerator, "/", final_denominator)
		} else if mixed_whole != 0 && mixed_whole < 0 {
			numerator_abs := abs(mixed_numerator)
			denominator_abs := abs(final_denominator)
			
			fmt.Println("Solution: ", mixed_whole, " ", numerator_abs, "/", denominator_abs)
		} else if mixed_whole == 0 && (final_denominator < 0 || final_numerator < 0) { 
			numerator_abs := abs(mixed_numerator)
			denominator_abs := abs(final_denominator)

			fmt.Println("Solution: ", numerator_abs * -1, "/", denominator_abs)
		} else {
			fmt.Println("Solution: ", final_numerator, "/", final_denominator)
		}
	default:
		fmt.Println("Invalid sign, please try again.")
		fractionSimplifier()
	} 
}

// Pre Alegbra/ Elementary Algebra functions
// For the average allow a maximum of 50 entries
func average() {
	// Initialize variables
	var entry_average float64
	var added_average float64
	var num_enteries int

	// Get the number of enteries
	fmt.Println("How many items would you like to average? Select a number between 1 and 50.")
	fmt.Scanln(&num_enteries)

	// Add all entries together
	if num_enteries > 50 || num_enteries <= 0 {
		fmt.Println("Invalid number of enteries, please try again.")
		average()
	} else { 
		for i := 0; i < num_enteries; i++ {
			fmt.Println("Enter the next entry.")
			fmt.Scanln(&entry_average)
			added_average += entry_average
		}
		// Divide added_average by num_enteries
		fmt.Println("The average is:", added_average / float64(num_enteries))
	}
}

func probability() {
	// Get user input for probability and number of trials
	var probability float64
	var num_trials int

	// Take input of probability?
	fmt.Println("What is the probability of success?")

	fmt.Scanln(&probability)

	// take input of trials
}

// func quadraticFormula() {

// }


/* Main will be used to call each section as necessary ranging from:
1. Pre-Algebra / Elementary Algebra
2. Intermediate Algebra / Coordinate Geometry 
3. Plane Geometry
4. Trigonometry */
func main() {
	var calc_choice int 
	// Select which function to use:
	fmt.Println("Which calculator would you like to use?")
	fmt.Println("1. Fraction Calculator")
	fmt.Println("2. Average Calculator")
	fmt.Println("3. Probability Calculator")
	fmt.Println("4. Quadratic Formula Calculator")
	fmt.Println("5. Exit")
	fmt.Scanln(&calc_choice)

	switch calc_choice { 
		case 1:
			fractionSimplifier()
		case 2:
			average()
		case 3:
			probability()
		case 5:
			fmt.Println("Exiting...")
		default:
			fmt.Println("Invalid choice, please try again.")
			main()
	}
}