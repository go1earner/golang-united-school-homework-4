package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	isDoubleNegative := false
	firstNumbIndex := 0
	secondNumbIndex := 1
	err = nil
	var splitted []string
	if strings.TrimSpace(input) == "" {
		return "", fmt.Errorf("input is empty %w", errorEmptyInput)
	} else if len(strings.TrimSpace(input)) <= 1 {
		return "", fmt.Errorf("expecting two operands, but received more or less %w", errorNotTwoOperands)
	} else if len(strings.TrimSpace(input)) > 4 {
		return "", fmt.Errorf("expecting two operands, but received more or less %w", errorNotTwoOperands)
	}
	if strings.Contains(input, "+") {
		splitted = strings.Split(input, "+")
		if len(splitted) != 2 {
			return "", fmt.Errorf("expecting two operands, but received more or less %w", errorNotTwoOperands)
		}
	} else if strings.Count(input, "-") >= 2 {
		splitted = strings.Split(input, "-")
		if splitted[0] == "" {
			splitted = splitted[1:]
		}
		if len(splitted) != 2 {
			return "", fmt.Errorf("expecting two operands, but received more or less %w", errorNotTwoOperands)
		}
		isDoubleNegative = true
	} else if strings.Count(input, "-") == 1 {
		splitted = strings.Split(strings.ReplaceAll(input, "-", "+"), "+")
		if len(splitted) != 2 {
			return "", fmt.Errorf("expecting two operands, but received more or less %w", errorNotTwoOperands)
		}
		firstNumber, err1 := strconv.Atoi(strings.TrimSpace(splitted[firstNumbIndex]))
		if err1 != nil {
			return "", fmt.Errorf("bad token. %w", err1)
		}
		secondNumber, err2 := strconv.Atoi(strings.TrimSpace(splitted[secondNumbIndex]))
		if err2 != nil {
			return "", fmt.Errorf("bad token. %w", err2)
		}
		output = strconv.Itoa(firstNumber - secondNumber)
		return
	}

	firstNumber, err1 := strconv.Atoi(strings.TrimSpace(splitted[firstNumbIndex]))
	if err1 != nil {
		return "", fmt.Errorf("bad token. %w", err1)
	}
	secondNumber, err2 := strconv.Atoi(strings.TrimSpace(splitted[secondNumbIndex]))
	if err2 != nil {
		return "", fmt.Errorf("bad token. %w", err2)
	}
	sum := firstNumber + secondNumber
	if isDoubleNegative {
		output = strconv.Itoa(-sum)
	} else {
		output = strconv.Itoa(sum)
	}
	return
}
