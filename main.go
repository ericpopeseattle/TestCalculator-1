package main

import (
	"TestCalculator/calculator"
	"fmt"
	"log"
	"os"
)

func main() {
	var argument1 string
	var argument2 string
	var operator string
	var result float64
	var err error

	argument1, argument2, operator, result, err = calculatorBody()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s %s %s = %f\n", argument1, operator, argument2, result)
}

func calculatorBody() (string, string, string, float64, error) {
	args := os.Args[1:]

	if len(args) != 3 {
		log.Fatal("Usage: main <operand1> <operator> <operand2>")
	}

	operand1 := parseOperand(args[0])
	operand2 := parseOperand(args[2])
	operator := args[1]

	result, err := calculator.Calculate(operand1, operand2, operator)

	return args[0], args[2], operator, result, err
}

func parseOperand(operand string) float64 {
	value, err := calculator.ParseOperand(operand)
	if err != nil {
		log.Fatal(err)
	}
	return value
}
