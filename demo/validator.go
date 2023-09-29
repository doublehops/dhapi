package main

import (
	"encoding/json"
	"fmt"
)

const (
	defaultErrorMessage = "field is not valid"
)

type ValidationFunctions func(bool, interface{}) (bool, string)

type Rule struct {
	VariableName string
	Value        interface{} // WHAT WAS THIS?
	Required     bool
	Function     []ValidationFunctions // AND WHAT WAS THIS?
}

type Error string

type ErrorMessages map[string][]Error

func RunValidation(rules []Rule) ErrorMessages {
	errorMessages := make(ErrorMessages)

	for _, prop := range rules {
		var errors []Error
		for _, rule := range prop.Function {
			valid, errMsg := rule(prop.Required, prop.Value)
			if !valid {
				errors = append(errors, Error(errMsg))
			}
		}

		if errors != nil {
			errorMessages[prop.VariableName] = errors
		}
	}

	return errorMessages
}

const MinLengthDefaultMessage = "is not the minimum length"

func MinLength(minLength int, errorMessage string) ValidationFunctions {
	return func(required bool, value interface{}) (bool, string) {
		if errorMessage == "" {
			errorMessage = MinLengthDefaultMessage
		}

		if v, ok := value.(string); ok {
			if required && v == "" {
				return true, ""
			}
			if len(v) >= minLength {
				return true, ""
			}
		}

		return false, errorMessage
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
		{"Make", carInput.Make, false, []ValidationFunctions{MinLength(13, "ErrorOverride")}},
	}

	errors := RunValidation(rules)
	j, _ := json.Marshal(errors)
	fmt.Println(string(j))
}
