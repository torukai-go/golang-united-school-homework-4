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

type MyError struct {
	Message string
}

func (e MyError) Error() string {
	return e.Message
}

func StringSum(input string) (output string, err error) {

	if len(strings.Fields(input)) == 0 {
		return "", fmt.Errorf("error:%w", errorEmptyInput)
	}

	var outputSlice = make([]string, 0, 0)
	var resultStr string

	output = strings.TrimSpace(input)
	output = strings.TrimPrefix(output, "+")

	if !strings.Contains(output, "-") {
		numbers := strings.Split(output, "+")
		a, err := strconv.Atoi(numbers[0])
		if err != nil {
			return "", fmt.Errorf("error:%w", err)
		}
		b, err := strconv.Atoi(numbers[1])
		if err != nil {
			return "", fmt.Errorf("error:%w", err)
		}
		resultStr = strconv.Itoa(a + b)
	} else {

		var num1Negative bool = false
		var num2Negative bool = false

		if strings.HasPrefix(output, "-") {
			output = strings.TrimPrefix(output, "-")
			num1Negative = true
		}

		if strings.Contains(output, "-") {
			outputSlice = strings.Split(output, "-")
			num2Negative = true
		} else {
			outputSlice = strings.Split(output, "+")
			num2Negative = false
		}
		a, _ := strconv.Atoi(string(outputSlice[0]))
		b, _ := strconv.Atoi(string(outputSlice[1]))

		if num1Negative {
			a = 0 - a
		}

		if num2Negative {
			b = 0 - b
		}

		resultStr = strconv.Itoa(a + b)
	}
	return resultStr, nil
}
