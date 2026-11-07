package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1            // unsigned integer
	z      complex128 = cmplx.Sqrt(-5 + 12i) // complex
)

func main() {

	// %T is use to print the type of the variable
	// %v is use to print the Value of the variable
	// %s is use to print the Value of the String
	// %q is use to print the Value of the String in Double or Single quotes

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

// GO Data Types

/*
 1. Basic -  int, float64, complex128, string, bool

 2. Aggregate - Array , Struct

 3. Refrence -  slice, map, function, channel, pointer

 4. Interface - interface, any
*/
