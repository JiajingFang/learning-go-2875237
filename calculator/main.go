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
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Value 1: ")
	input1, _ := reader.ReadString('\n')
	float1 := valiValue(input1)

	fmt.Print("Value 2: ")
	input2, _ := reader.ReadString('\n')
	float2 := valiValue(input2)

	fmt.Print("Select your operation +-*/: ")
	input3, _ := reader.ReadString('\n')
	input3 = strings.TrimSpace(input3)

	switch input3 {
	case "+":
		add(float1, float2)

	case "-":
		minus(float1, float2)

	case "*":
		multi(float1, float2)

	case "/":
		division(float1, float2)

	}

}

func valiValue(input string) float64 {
	float1, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value must be a number")
	}
	return float1
}

func add(value1, value2 float64) {
	sum := value1 + value2
	sum = math.Round(sum*100) / 100
	fmt.Printf("The sum of %v and %v is %v\n\n", value1, value2, sum)
}

func minus(value1, value2 float64) {
	sum := value1 - value2
	sum = math.Round(sum*100) / 100
	fmt.Printf("The minus of %v and %v is %v\n\n", value1, value2, sum)
}

func multi(value1, value2 float64) {
	sum := value1 * value2
	sum = math.Round(sum*100) / 100
	fmt.Printf("The multiplication of %v and %v is %v\n\n", value1, value2, sum)
}

func division(value1, value2 float64) {
	sum := value1 / value2
	sum = math.Round(sum*100) / 100
	fmt.Printf("The division of %v and %v is %v\n\n", value1, value2, sum)
}
