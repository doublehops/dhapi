package main

import (
	"encoding/json"
	"fmt"
)

const (
	defaultErrorMessage = "field is not valid"
)

type ValidationFunctions func(interface{}) (bool, string)

type Rule struct {
	VariableName string
	Value        interface{} // WHAT WAS THIS?
	Required     bool
	Function     []ValidationFunctions // AND WHAT WAS THIS?
}

type VariableName string
type Error string

type ErrorMessages map[string][]Error

//type ErrorMessages struct {
//	VariableName string  `json:"name"`
//	Errors       []Error `json:"errors"`
//}

func RunValidation(rules []Rule) ErrorMessages {
	errorMessages := make(ErrorMessages)

	for _, prop := range rules {
		var errors []Error
		for _, rule := range prop.Function {
			valid, errMsg := rule(prop.Value)
			if !valid {
				errors = append(errors, Error(errMsg))
			}
		}

		if errors != nil {
			errorMessages[prop.VariableName] = errors
			//errorMessages = append(errorMessages, ErrorMessages{
			//	VariableName: prop.VariableName,
			//	Errors:       errors,
			//})
		}
	}

	return errorMessages
}

const MinLengthError = "is not the minimum length"

func MinLength(minLength int) ValidationFunctions {
	return func(value interface{}) (bool, string) {
		if v, ok := value.(string); ok {
			if len(v) == 0 && v == "" {
				return false, MinLengthError
			}
			if len(v) >= minLength {
				return true, ""
			}
		}
		return false, MinLengthError
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
		{"Make", carInput.Make, true, []ValidationFunctions{MinLength(13)}},
	}

	errors := RunValidation(rules)
	j, _ := json.Marshal(errors)
	fmt.Println(string(j))
}
