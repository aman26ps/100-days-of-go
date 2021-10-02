package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	fmt.Println(os.Args[1])

	//run the program as "go run  day-1.go lets-start-100-days-of-code"
	// and it will take the above argument and print it on the terminal

	//path package

	// the Split package will bifarcate the folder and the files
	// and return both of the values as a string

	dir, file := filepath.Split("marvel/spiderman.exe")

	// join will formulate the different folders in a workspace into a
	// folder structured path for example in below case it will be
	// a/b/c --> like a folder structure
	joinPath := filepath.Join("a", "/b", "c")

	fmt.Println("dir and filename are ", dir, file)
	fmt.Println(joinPath)
	fmt.Println(filepath.Match("a*", "aman"))

	//type conversion

	distance := 1000

	time := 2.4
	//by default the time will be converted into int as
	// we cant divide int with float64 type but that will
	//give us inaccurate result, thats why how to perform
	//type conversion is very important
	speed := int(float64(distance) / time)

	fmt.Println(speed)

}
