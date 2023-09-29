package validator

import (
	"testing"
)

func TestMinLength(t *testing.T) {
	tests := []struct {
		name           string
		value          string
		required       bool
		length         int
		expectedResult bool
		expectedError  string
	}{
		{
			name:           "minLengthPass",
			value:          "orange",
			required:       true,
			length:         6,
			expectedResult: true,
			expectedError:  "",
		},
		{
			name:           "minLengthEmptyButNotRequired",
			value:          "",
			required:       false,
			length:         6,
			expectedResult: true,
			expectedError:  "",
		},
		{
			name:           "minLengthEmptyAndRequired",
			value:          "",
			required:       true,
			length:         6,
			expectedResult: false,
			expectedError:  MinLengthDefaultMessage,
		},
		{
			name:           "tooShortAndRequired",
			value:          "apple",
			required:       false,
			length:         6,
			expectedResult: false,
			expectedError:  MinLengthDefaultMessage,
		},
		{
			name:           "testCustomErrorMessage",
			value:          "apple",
			required:       false,
			length:         6,
			expectedResult: false,
			expectedError:  "myCustomMessage",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := MinLength(tt.length, tt.expectedError)
			res, err := validator(tt.required, tt.value)
			if tt.expectedResult != res {
				t.Errorf("test result not as expected. Expected: %t; got: %t", tt.expectedResult, res)
			}
			if tt.expectedError != err {
				t.Errorf("test result not as expected. Expected: %s; got: %s", tt.expectedError, err)
			}
		})
	}
}
