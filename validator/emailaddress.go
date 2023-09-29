package validator

import "regexp"

const (
	EmailAddressDefaultMessage = "is not a valid email address"

	emailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)

func EmailAddress(errorMessage string) ValidationFunctions {
	return func(required bool, value interface{}) (bool, string) {
		if errorMessage == "" {
			errorMessage = EmailAddressDefaultMessage
		}

		if v, ok := value.(string); ok {
			if v == "" && !required {
				return true, ""
			}

			re := regexp.MustCompile(emailPattern)
			if re.MatchString(v) {
				return true, ""
			}
		}

		return false, errorMessage
	}
}
