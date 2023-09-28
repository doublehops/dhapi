package main

import "fmt"

const (
	defaultErrorMessage = "field is not valid"
)

type ValidationFunction func(interface{}) bool

type Rule struct {
	VariableName string
	Value        interface{} // WHAT WAS THIS?
	Required     bool
	Function     ValidationFunction // AND WHAT WAS THIS?
}

type ErrorMessage struct {
	VariableName string
	Errors       string
}

func RunValidation(rules []Rule) []ErrorMessage {
	var errors []ErrorMessage

	for _, rule := range rules {
		valid := rule.Function(rule.Value)
		if !valid {
			errors = append(errors, ErrorMessage{rule.VariableName, defaultErrorMessage})
		}
	}

	return errors
}

func MinLength(minLength int) ValidationFunction {
	return func(value interface{}) bool {
		if v, ok := value.(string); ok {
			if len(v) == 0 && v == "" {
				return false
			}
			if len(v) >= minLength {
				return true
			}
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
		{"Make", carInput.Make, true, MinLength(13)},
	}

	errors := RunValidation(rules)
	fmt.Printf("%v", errors)
}
