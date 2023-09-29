package validator

const (
	RequiredPropertyError   = "this is a required property"
	ProcessingPropertyError = "unable to process property"
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
