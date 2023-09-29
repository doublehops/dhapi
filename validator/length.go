package validator

const (
	MinLengthDefaultMessage     = "does not meet minimum length"
	MaxLengthDefaultMessage     = "exceeds maximum length"
	BetweenLengthDefaultMessage = "does not conform to min and max lengths"
)

func MinLength(minLength int, errorMessage string) ValidationFunctions {
	return func(required bool, value interface{}) (bool, string) {
		if errorMessage == "" {
			errorMessage = MinLengthDefaultMessage
		}

		if v, ok := value.(string); ok {
			if v == "" && !required {
				return true, ""
			}
			if len(v) >= minLength {
				return true, ""
			}
		}

		return false, errorMessage
	}
}

func MaxLength(maxLength int, errorMessage string) ValidationFunctions {
	return func(required bool, value interface{}) (bool, string) {
		if errorMessage == "" {
			errorMessage = MaxLengthDefaultMessage
		}

		if v, ok := value.(string); ok {
			if required && v == "" {
				return false, errorMessage
			}
			if len(v) <= maxLength {
				return true, ""
			}
		}

		return false, errorMessage
	}
}

func Between(minLength, maxLength int, errorMessage string) ValidationFunctions {
	return func(required bool, value interface{}) (bool, string) {
		if errorMessage == "" {
			errorMessage = BetweenLengthDefaultMessage
		}

		if v, ok := value.(string); ok {
			if !required && v == "" {
				return true, ""
			}
			if len(v) >= minLength && len(v) <= maxLength {
				return true, ""
			}
		}

		return false, errorMessage
	}
}
