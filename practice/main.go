package main

import (
	"fmt"
)

const aConst int = 64

func main() {

	var aString string = "hello from gooooo"

	fmt.Println(aString)
	fmt.Printf("the variable's type is %T\n", aString)

	var anInt int = 42
	fmt.Println(anInt)

	var defInt int
	fmt.Println(defInt)

	var inInt = 50
	fmt.Println(inInt)

	newInt := 30
	fmt.Println(newInt)

	fmt.Println("print the const", aConst)
}
