package main

import (
	"fmt"
	"os"
	"path"
)

func main() {

	//variable assignment

	color := "green"

	// if we use color := "blue" we will face NonewVar error thrown by compiler
	// because short variable declaration should declare new variable in our case we are using
	// color variable to overwrite so we will use "=" operand only.
	//color = "blue"

	color = "dark " + color
	fmt.Println(color) //prints blue

	n := 0.

	n = 3.14 * 2

	fmt.Println(n)

	var (
		perimeter     int
		width, height = 5, 6
	)

	perimeter = 2 * (width + height)
	fmt.Println(perimeter)

	var (
		lang    string
		version int
	)

	lang = "go"
	version = 2

	// DO NOT TOUCH THIS
	fmt.Println(lang, "version", version)

	var (
		planet string
		isTrue bool
		temp   float64
	)

	planet = "mars"
	temp = 19.5
	isTrue = true

	fmt.Println("Air is good on", planet)
	fmt.Println("It's", isTrue)
	fmt.Println("It is", temp, "degrees")

	//discarding the return value of multi function below
	//and printing only the second integer return value

	_, b := multi()

	fmt.Println(b)

	color, color2 := "red", "blue"

	color, color2 = "orange", "green"

	fmt.Println(color, color2)

	red, blue := "red", "blue"

	red, blue = "blue", "red"

	fmt.Println(red, blue)

	a, _ := path.Split("secret/file.txt")

	fmt.Println(a)

	//fix the code
	// a, b := 10, 5.5
	// fmt.Println(a + b)

	e, f := 10, 5.5

	fmt.Println(float64(e) + f)

	// a, b := 10, 5.5
	// a = b
	// fmt.Println(a + b)

	g, h := 10, 5.5
	g = int(h)

	fmt.Println(float64(g) + h)

	fmt.Println(float64(5.5))

	age := 2
	fmt.Println(float64(7.5) + float64(age))

	min := int8(127)
	max := int16(1000)

	// FIX THE CODE HERE
	fmt.Println(max + int16(min))

	// UNCOMMENT & FIX THIS CODE
	// count := ?

	//count := len(os.Args[1] + os.Args[2] + os.Args[3])

	// UNCOMMENT IT & THEN DO NOT TOUCH THIS CODE
	//fmt.Printf("There are %d names.\n", count)

	name := os.Args[1]

	fmt.Println(name)

}

// multi is a function that returns multiple int values
func multi() (int, int) {
	return 5, 4
}
