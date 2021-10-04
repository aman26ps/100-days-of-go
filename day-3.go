package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {

	//string literal and raw string literal

	a := "i am a string literal" //string literal is under double quotes and it is interpreted by go at run time

	b := ` i am a raw string literal
	s
	s
	s` // raw string literal never get interpreted by go at run time and the characters inside the `` will be ignored and presented as it is

	fmt.Println(a)
	fmt.Println(b)

	//the below code will find the bytes used in the string as well as the characters

	str := "İNANÇ"
	fmt.Println("bytes =", len(str))
	fmt.Println("runes =", utf8.RuneCountInString(str))

	//excercise take the argument from the console and print the argument in caps with exclamation
	// in front and at end of the string for example, hey should be printed as !!!HEY!!!

	m := os.Args[1]
	m = strings.ToLower(m)

	//the below utf8 package helps to count special variables input charactwers

	lens := utf8.RuneCountInString(m)

	// string.repeat will repeat the characters as many times as lens return value

	fmt.Println(strings.Repeat("!", lens) + m + strings.Repeat("!", lens))

	path := `c:\program files\duper super\fun.txt
c:\program files\really\funny.png`
	fmt.Println(path)

	json := `

{
        "Items": [{
                "Item": {
                        "name": "Teddy Bear"
                }
        }]
}	
`

	fmt.Println(json)

	msg := `hi ` + m + `!` + `
how are you?`

	fmt.Println(msg)

	msgs := `
	
	The weather looks good.
I should go and play.
	`

	//below will trim spaces from above message

	fmt.Println(strings.TrimSpace(msgs))

	//below wil remove the right spaces from the string

	as := "sas      "
	fmt.Println(len(strings.TrimRight(as, " ")))

}
