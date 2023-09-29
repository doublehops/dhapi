package validator

const (
	MinValueDefaultMessage = "is below required amount"
	MaxValueDefaultMessage = "is above required amount"
	InRangeDefaultMessage  = "is not within required range"
)

func MinValue(minValue int, errorMessage string) ValidateFuncs {
	return func(required bool, value interface{}) (bool, string) {
		if errorMessage == "" {
			errorMessage = MinValueDefaultMessage
		}

		var (
			v  int
			ok bool
		)

		if v, ok = value.(int); !ok {
			return false, ProcessingPropertyError
		}

		if v == 0 && !required {
			return true, ""
		}

		if v < minValue {
			return false, errorMessage
		}

		return true, ""
	}
}

func MaxValue(maxValue int, errorMessage string) ValidateFuncs {
	return func(required bool, value interface{}) (bool, string) {
		if errorMessage == "" {
			errorMessage = MaxValueDefaultMessage
		}

		var (
			v  int
			ok bool
		)

		if v, ok = value.(int); !ok {
			return false, ProcessingPropertyError
		}

		if v > maxValue {
			return false, errorMessage
		}

		return true, ""
	}
}

func IntInRange(minValue, maxValue int, errorMessage string) ValidateFuncs {
	return func(required bool, value interface{}) (bool, string) {
		if errorMessage == "" {
			errorMessage = InRangeDefaultMessage
		}

		var (
			v  int
			ok bool
		)

		if v, ok = value.(int); !ok {
			return false, ProcessingPropertyError
		}

		if v == 0 && !required {
			return true, ""
		}

		if v < minValue || v > maxValue {
			return false, errorMessage
		}

		return true, ""
	}
}
