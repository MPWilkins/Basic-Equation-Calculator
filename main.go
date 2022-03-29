package main

import (
	"fmt"
	//"math"
)

// Documentation help: https://pkg.go.dev/fmt

// Basic fraction simplifier formula, something along lines of a/b +/- cd
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
	fmt.Print("Addition(+) or Subtraction(-): ")
	fmt.Scanf("%s", &sign)

	switch sign {
	case "-":
		a = a * d
		c = c * b
		e = b * d
		
		f = a - c

		fmt.Printf("Before Simplifying the answer is: %d/%d", f, e)

	case "+":
		a = a * d
		c = c * b
		e = b * d

	} 

	// fmt.Println(e)
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
}