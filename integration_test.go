//go:build integration
// +build integration

package main

import (
	"errors"
	"os"
	"strconv"
	"testing"
)

func TestMainFunct(t *testing.T) {
	var argument1 string
	var argument2 string
	var operator string
	var result float64
	var err error

	tests := []struct {
		operand1 float64
		operator string
		operand2 float64
		expected float64
		err      error
	}{
		{1, "+", 1, 2, nil},
		{5, "-", 3, 2, nil},
		{2, "*", 3, 6, nil},
		{6, "/", 2, 3, nil},
		{6, "/", 0, 0, errors.New("division by zero is not allowed")},
		{1, "%", 1, 0, errors.New("invalid operator")},
	}

	for _, test := range tests {
		argument := make([]string, len(os.Args))
		argument[0] = os.Args[0]
		argument[1] = strconv.FormatFloat(test.operand1, 'f', -1, 64)
		argument[2] = test.operator
		argument[3] = strconv.FormatFloat(test.operand2, 'f', -1, 64)
		os.Args = argument

		argument1, argument2, operator, result, err = calculatorBody()

		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("Expected error: %v, got: %v", test.err, err)
		}

		if result != test.expected {
			t.Errorf("From Test %s %s %s, Expected result: %v, got: %v", argument1, operator, argument2, test.expected, result)
		}
	}
}
