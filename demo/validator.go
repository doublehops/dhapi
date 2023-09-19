package main

import "fmt"

const (
	defaultErrorMessage = "field is not valid"
)

type Rule struct {
	VariableName string
	Value        any // could be any type.
	Required     bool
	Function     func() bool // could have any signature.
}

type ErrorMessage struct {
	VariableName string
	Errors       string
}

func RunValidation(rules []Rule) []ErrorMessage {
	var errors []ErrorMessage

	for _, rule := range rules {
		err := rule.Function
		if err != nil {
			errors = append(errors, ErrorMessage{rule.VariableName, defaultErrorMessage})
		}
	}

	return errors
}

func MinLength() func(value string, minLength int) bool {
	return func(value string, minLength int) bool {
		if len(value) == 0 && value == "" {
			return true
		}

		if len(value) >= minLength {
			return true
		}

		return false
	}
}

/****************  USAGE  ******************/

type Car struct {
	Make  string
	Model string
	Year  int
}

func main() {

	carInput := Car{
		Make:  "Ford",
		Model: "Falcon",
		Year:  1972,
	}

	rules := []Rule{
		{"Make", &carInput.Make, true, MinLength},
	}

	errors := RunValidation(rules)
	fmt.Printf("%v", errors)
}
