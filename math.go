package main

import (
	"fmt"
	"math"
)

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

// To do later: Number of Decimal places for a solution
// Look at the Probability Calculator for the proof of concept
/* func round(num float64) float64 {
	num_dec_places float64
	return math.Round(num * num_dec) / 100000
} */

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

// Pre-Algebra / Elementary Algebra
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
	var num_trials float64
	var dec_places float64

	// Take input of probability?
	fmt.Println("Chances that given event occurs? (i.e. rolling the number 3 on a 6 sided die)")
	fmt.Scanln(&probability)

	// take input of trials
	fmt.Println("How many possible outcomes for the event? (i.e. 6 sided die)")
	fmt.Scanln(&num_trials)

	fmt.Println("How many decimal places would you like to go to?")
	fmt.Scanln(&dec_places)
	
	// Calculate the probability
	prob_percentage := probability * 100
	sol_probability := prob_percentage / math.Floor(num_trials)

	// Use Exponents to calculate the number of decimal places
	num_dec_places := math.Pow(10, math.Floor(dec_places))
	
	// Display the probability
	fmt.Println("The probability is:", math.Round(sol_probability*num_dec_places)/num_dec_places, "%")
}

// x= (-b±√(b²-4ac))/(2a)
func quadraticFormula() {
	// Get user input for a, b, and c
	var a float64
	var b float64
	var c float64

	// Show equation layout.
	fmt.Println("x = (-b±√(b²-4ac))/(2a)")

	// Take input for a
	fmt.Println("Enter the value of a.")
	fmt.Scanln(&a)
	// Take input for b
	fmt.Println("Enter the value of b.")
	fmt.Scanln(&b)
	// Take input for c
	fmt.Println("Enter the value of c.")
	fmt.Scanln(&c)

	// For obtaining decimal places
	dec_places := math.Pow(10, 5)

	// Solve b²-4ac
	d := math.Pow(b, 2) - (4 * a * c)
	//check what d is
	fmt.Println("d is:", d)

	// Check if d is negative or not
	if d < 0 {
		// alter the value of d to be positive
		d *= -1
		
		// Calculate the square root of d
		sqrt_d := math.Sqrt(d)

		// Calculate first part of top half
		top_half_b_prep := b * -1

		// Calculate the bottom half
		bottom_half := 2 * a

		// Get Values of the combined values
		first_solution := top_half_b_prep / bottom_half
		second_solution := sqrt_d / bottom_half

		// Display answers with an i after the second solution
		fmt.Println("The Solutions are:")
		fmt.Println("x =", math.Round(first_solution*dec_places)/dec_places, "+", math.Round(second_solution*dec_places)/dec_places, "i")
		fmt.Println("x =", math.Round(first_solution*dec_places)/dec_places, "-", math.Round(second_solution*dec_places)/dec_places, "i")


	} else {
		// Get the square root of d if not a negative
	sqrt_d := math.Sqrt(d)
	// check sqrt_d
	fmt.Println("sqrt_d is:", sqrt_d)

	// Find the total of both halfs
	// Top half will have two different components due to two different answers
	top_half_add := (-b + sqrt_d)
	top_half_sub := (-b - sqrt_d)
	bot_half := (2 * a)
	// check top and bottom half
	fmt.Println("top_half is:", top_half_add)
	fmt.Println("top_half is:", top_half_sub)
	fmt.Println("bot_half is:", bot_half)

	x := top_half_add / bot_half
	y := top_half_sub / bot_half

	// Display the solution
	fmt.Println("The solution is for additon is:", math.Round(x*dec_places)/dec_places)
	fmt.Println("The solution is for subtraction is:", math.Round(y*dec_places)/dec_places)
	}
}

func basic_Algebra() {
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
		case 4:
			quadraticFormula()
		case 5:
			fmt.Println("Exiting...")
		default:
			fmt.Println("Invalid choice, please try again.")
			basic_Algebra()
	}
}

// Intermediate Algebra / Coordinate Geometry
// Distance Calculator - d=√(x₁ – x₂)² + (y₁ – y₂)²
func distance(){
	// Get user input
	var x1 float64
	var y1 float64
	var x2 float64
	var y2 float64

	// Show equation layout.
	fmt.Println("d=√(x₁ – x₂)² + (y₁ – y₂)²")

	fmt.Println("(x₁, y₁), (x₂, y₂)")
	fmt.Println("Enter the value of x₁.")
	fmt.Scanln(&x1)
	fmt.Println("Enter the value of y₁.")
	fmt.Scanln(&y1)
	fmt.Println("Enter the value of x₂.")
	fmt.Scanln(&x2)
	fmt.Println("Enter the value of y₂.")
	fmt.Scanln(&y2)

	// Calculate the distance
	distance := math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))

	round_distance := math.Round(distance * 100000) / 100000

	fmt.Println("The distance is:", round_distance)
}

// Slope Calculator - y₂ – y₁ /  x₂ – x₁
func slope(){
	// Get user input
	var x1 float64
	var y1 float64
	var x2 float64
	var y2 float64

	// Show equation layout.
	fmt.Println("y₂ – y₁ /  x₂ – x₁")

	fmt.Println("(x₁, y₁), (x₂, y₂)")
	fmt.Println("Enter the value of x₁.")
	fmt.Scanln(&x1)
	fmt.Println("Enter the value of y₁.")
	fmt.Scanln(&y1)
	fmt.Println("Enter the value of x₂.")
	fmt.Scanln(&x2)
	fmt.Println("Enter the value of y₂.")
	fmt.Scanln(&y2)

	// Calculate the slope
	y_slope := y2 - y1
	x_slope := x2 - x1

	fmt.Println("The slope is:", y_slope, "/", x_slope)

	// Reduce slope to to a decimal
	slope := y_slope / x_slope

	// make slope only have five decimal places
	round_slope := math.Round(slope * 100000) / 100000
	fmt.Println("The final slope when simplified is (Rounded to 5 decimal places):", round_slope)
}

// Slope Intercept Calculator - y = mx + b
func slope_intercept(){
	// Get user input
	var x1 float64
	var y1 float64
	var x2 float64
	var y2 float64

	// Show equation layout.
	fmt.Println("y = mx + b")

	fmt.Println("(x₁, y₁), (x₂, y₂)")
	fmt.Println("Enter the value of x₁.")
	fmt.Scanln(&x1)
	fmt.Println("Enter the value of y₁.")
	fmt.Scanln(&y1)
	fmt.Println("Enter the value of x₂.")
	fmt.Scanln(&x2)
	fmt.Println("Enter the value of y₂.")
	fmt.Scanln(&y2)

	// Calculate the slope - m
	y_slope := y2 - y1
	x_slope := x2 - x1

	fmt.Println("The slope is:", y_slope, "/", x_slope)

	// Reduce slope to to a decimal
	m := y_slope / x_slope
}

func algebra_Geo(){
	var calc_choice int

	// Select a calculator:
	fmt.Println("Which calculator would you like to use?")
	fmt.Println("1. Distance Calculator")
	fmt.Println("2. Slope Calculator")
	fmt.Println("3. Slope Intercept Calculator")
	fmt.Println("4. Midpoint Calculator")
	fmt.Println("5. Exit")

	fmt.Scanln(&calc_choice)

	switch calc_choice {
		case 1:
			distance()
		case 2:
			slope()
		case 3:
		case 4:
		case 5:
			fmt.Println("Exiting...")
	}
}

/* Main will be used to call each section as necessary ranging from:
1. Pre-Algebra / Elementary Algebra
2. Intermediate Algebra / Coordinate Geometry 
3. Plane Geometry
4. Trigonometry 

https://www.collegeraptor.com/getting-in/articles/act-sat/5-critical-formulas-you-must-know-for-the-act/ */

func main() {
	//Group of calculators
	var calc_choice int


	// Select a calculator:
	fmt.Println("Which calculator would you like to use?")
	fmt.Println("1. Pre-Algebra / Elementary Algebra")
	fmt.Println("2. Intermediate Algebra / Coordinate Geometry")
	fmt.Println("5. Exit")

	fmt.Scanln(&calc_choice)

	switch calc_choice {
		case 1:
			basic_Algebra()
		case 2:
			algebra_Geo()
		case 5:
			fmt.Println("Exiting...")
	}

}