package main

import (
	"fmt"
	//"math"
)

// Documentation help: https://pkg.go.dev/fmt

// Basic fraction simplifier formula, something along lines of a/b +/- cd
// If the denomiator is 1 as the end of the math solution, then remove the denominator after it is simplified.
// Greatest Common Denominator
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}


func fractionSimplifier() {
	var a, b, c, d, e, f int
	var sign string

	fmt.Println("Enter numbers for the following format: A/B +/- C/D")
	fmt.Print("A: ")
	fmt.Scanf("%d", &a)
	fmt.Print("B: ")
	fmt.Scanf("%d", &b)
	fmt.Print("C: ")
	fmt.Scanf("%d", &c)
	fmt.Print("D: ")
	fmt.Scanf("%d", &d)
	fmt.Print("Addition(+) or Subtraction(-). Use the respective sign: ")
	fmt.Scanf("%s", &sign)

	// 1. Find a method to find the Great Common Denominator
	// search for proper contol input type validation.
	
	switch sign {
	case "-":
		a = a * d
		c = c * b
		e = b * d
		
		f = a - c

		x := gcd(f, e)


		fmt.Printf("Before Simplifying the answer is: %d/%d", f, e)
		fmt.Printf("\nAfter Simplifying the answer is: %d/%d", f/x, e/x)

	case "+":
		a = a * d
		c = c * b
		e = b * d

		f = a + c

		x := gcd(f, e)


		fmt.Printf("Before Simplifying the answer is: %d/%d", f, e)
		fmt.Printf("\nAfter Simplifying the answer is: %d/%d", f/x, e/x)

	default:
		fmt.Println("Invalid Input. Please use a + or - sign.")
	} 
}

// Pre Alegbra/ Elementary Algebra functions
// Will use %g for floats
func average() {

}

func probability() {

}

func quadraticFormula() {

}


/* Main will be used to call each section as necessary ranging from:
1. Pre-Algebra / Elementary Algebra
2. Intermediate Algebra / Coordinate Geometry 
3. Plane Geometry
4. Trigonometry */
func main() {
	fractionSimplifier()
	// x := gcd(3, 9)
	// fmt.Println(x)
}