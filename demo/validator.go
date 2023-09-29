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

type Person struct {
	Name         string
	Age          string
	EmailAddress string
}

func main() {

	person := Person{
		Name:         "Jo",
		Age:          "Smith",
		EmailAddress: "jo.smith",
	}

	rules := []Rule{
		{"name", person.Name, true, []ValidationFunctions{MinLength(13, "")}},
		{"emailAddress", person.EmailAddress, false, []ValidationFunctions{EmailAddress("")}},
	}

	errors := RunValidation(rules)
	j, _ := json.Marshal(errors)
	fmt.Println(string(j))
}
