//go:build unit
// +build unit

package test

import (
	"TestCalculator/calculator"
	"errors"
	"testing"
)

func TestCalculate_edge(t *testing.T) {
	tests := []struct {
		operand1 float64
		operator string
		operand2 float64
		expected float64
		err      error
	}{
		{10000, "+", 0.00001, 10000.00001, nil},
		{0.00001, "+", 0.00001, 0.00002, nil},
		{2.0002, "*", 3, 6.0006, nil},
		{13.716, "/", 4.572, 3, nil},
		{0, "/", 6, 0, errors.New("division by zero is not allowed")},
		{1, "a", 1, 0, errors.New("invalid operator")},
		{1, "^", 1, 0, errors.New("invalid operator")},
		{1, "++++++++", 1, 0, errors.New("invalid operator")},
	}

	for _, test := range tests {
		result, err := calculator.Calculate(test.operand1, test.operand2, test.operator)

		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("Expected error: %v, got: %v", test.err, err)
		}

		if result != test.expected {
			t.Errorf("Expected result: %v, got: %v", test.expected, result)
		}
	}
}

func TestParseOperand_edge(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
		err      error
	}{
		{"1", 1, nil},
		{"3.14", 3.14, nil},
		{"abc", 0, errors.New("invalid operand")},
	}

	for _, test := range tests {
		result, err := calculator.ParseOperand(test.input)

		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("Expected error: %v, got: %v", test.err, err)
		}

		if result != test.expected {
			t.Errorf("Expected result: %v, got: %v", test.expected, result)
		}
	}
}
