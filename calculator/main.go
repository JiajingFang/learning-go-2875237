package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("A simple calculator")

	fmt.Print("Value 1: ")
	reader := bufio.NewReader(os.Stdin)
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if err != nil {
		panic("value 1 must be a number")
	}

	fmt.Print("Value 2: ")
	reader = bufio.NewReader(os.Stdin)
	numInput, _ = reader.ReadString('\n')
	bFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if err != nil {
		panic("value 2 must be a number")
	}

	fmt.Println("the sum of ", aFloat, " and ", bFloat, " is ", math.Round((aFloat+bFloat)*100)/100)
}
