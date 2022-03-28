package main

import (
	"fmt"
	//"math"
)

// Basic fraction simplifier formula, something along lines of a/b +/- cd
func fractionSimplifier() {
	var a, b, c, d int
	var sign string

	fmt.Println("Enter numbers for the following format: A/B +/- C/D")
	fmt.Print("A: ")
	fmt.Scan(&a)
	fmt.Print("B: ")
	fmt.Scan(&b)
	fmt.Print("C: ")
	fmt.Scan(&c)
	fmt.Print("D: ")
	fmt.Scan(&d)
	fmt.Print("Addition or Subtraction: ")
	fmt.Scan(&sign)
}

// Pre Alegbra/ Elementary Algebra functions
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