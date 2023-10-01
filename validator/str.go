package validator

const (
	keyNotInSlice = "value was not found in available items"
)

func In(slice []interface{}, errorMessage string) ValidateFuncs {
	return func(required bool, value interface{}) (bool, string) {
		if errorMessage == "" {
			errorMessage = keyNotInSlice
		}

		// We want to convert the slice values into strings.
		var strSlice []string
		for _, item := range slice {
			strSlice = append(strSlice, item.(string))
		}

		v := value.(string)

		if v == "" && required {
			return false, RequiredPropertyError
		}

		if v == "" && !required {
			return true, ""
		}

		for _, item := range strSlice {
			if item == v {
				return true, ""
			}
		}

		return false, errorMessage
	}
}
