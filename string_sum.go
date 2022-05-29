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

	trimmed := strings.TrimSpace(input)

	if len(trimmed) == 0 {
		return "", fmt.Errorf("empty input error: %w", errorEmptyInput)
	}

	expression := ""

	signIndex := 0

	for i := 0; i < len(trimmed); i++ {

		if string(trimmed[i]) == " " {
			continue
		}

		if string(trimmed[i]) == "+" {
			continue
		}

		if string(trimmed[i]) == "-" {
			expression += "-"
			continue
		}

		_, err := strconv.Atoi(string(trimmed[i]))
		if err != nil {
			return "", fmt.Errorf("character is not a digit: %w", err)
		}

		expression += string(trimmed[i])
	}

	for i := 1; i < len(expression); i++ {
		if string(expression[i]) == "+" || string(expression[i]) == "-" {
			signIndex = i
			break
		}
	}

	first := expression[:signIndex]
	second := expression[signIndex:]

	// if len(numbers) > 2 {
	// 	return "", fmt.Errorf("operand count missmatch: %w", errorNotTwoOperands)
	// }

	f, err := strconv.Atoi(first)
	if err != nil {
		return "", fmt.Errorf("character is not a digit: %w", err)
	}

	s, err := strconv.Atoi(second)
	if err != nil {
		return "", fmt.Errorf("character is not a digit: %w", err)
	}

	sum := f + s

	return strconv.Itoa(sum), nil
}
