package validator

const MinLengthDefaultMessage = "is not the minimum length"

func MinLength(minLength int, errorMessage string) ValidationFunctions {
	return func(required bool, value interface{}) (bool, string) {
		if errorMessage == "" {
			errorMessage = MinLengthDefaultMessage
		}

		if v, ok := value.(string); ok {
			if !required && v == "" {
				return true, ""
			}
			if len(v) >= minLength {
				return true, ""
			}
		}

		return false, errorMessage
	}
}
