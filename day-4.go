package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Iota Months
//
//  1. Initialize the constants using iota.
//  2. You should find the correct formula for iota.
//
// RESTRICTIONS
//  1. Remove the initializer values from all constants.
//  2. Then use iota once for initializing one of the
//     constants.
//
// EXPECTED OUTPUT
//  9 10 11
// ---------------------------------------------------------

func main() {
	const (
		Nov = -(iota - 11)
		Oct
		Sep
	)

	fmt.Println(Sep, Oct, Nov)

	const _ = iota
	//    ^- this is just a name

	// Now, use iota and initialize the following constants
	// "automatically" to 1,  ,2, and 3 respectively.
	const (
		Jan = iota + 1
		Feb
		Mar
	)

	// This should print: 1 2 3
	// Not: 0 1 2
	fmt.Println(Jan, Feb, Mar)

	const (
		Winter = iota + 12
		Spring
		Summer
		Fall
	)

	fmt.Println(Winter, Spring, Summer, Fall)
}
