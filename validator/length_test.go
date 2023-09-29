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

func TestMaxLength(t *testing.T) {
	tests := []struct {
		name           string
		value          string
		required       bool
		length         int
		expectedResult bool
		expectedError  string
	}{
		{
			name:           "maxLengthPass",
			value:          "orange",
			required:       true,
			length:         6,
			expectedResult: true,
			expectedError:  "",
		},
		{
			name:           "maxLengthEmptyButNotRequired",
			value:          "",
			required:       false,
			length:         6,
			expectedResult: true,
			expectedError:  "",
		},
		{
			name:           "maxLengthEmptyAndRequired",
			value:          "",
			required:       true,
			length:         6,
			expectedResult: false,
			expectedError:  MaxLengthDefaultMessage,
		},
		{
			name:           "tooLongAndRequired",
			value:          "apple",
			required:       true,
			length:         4,
			expectedResult: false,
			expectedError:  MaxLengthDefaultMessage,
		},
		{
			name:           "tooLongAndNotRequired",
			value:          "apple",
			required:       false,
			length:         4,
			expectedResult: false,
			expectedError:  MaxLengthDefaultMessage,
		},
		{
			name:           "testCustomErrorMessage",
			value:          "apple",
			required:       false,
			length:         4,
			expectedResult: false,
			expectedError:  "myCustomMessage",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := MaxLength(tt.length, tt.expectedError)
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

func TestBetween(t *testing.T) {
	tests := []struct {
		name           string
		value          string
		required       bool
		minLength      int
		maxLength      int
		expectedResult bool
		expectedError  string
	}{
		{
			name:           "betweenLengthPass",
			value:          "orange",
			required:       true,
			minLength:      3,
			maxLength:      6,
			expectedResult: true,
			expectedError:  "",
		},
		{
			name:           "betweenEmptyButNotRequired",
			value:          "",
			required:       false,
			minLength:      3,
			maxLength:      6,
			expectedResult: true,
			expectedError:  "",
		},
		{
			name:           "betweenTooShortAndRequired",
			value:          "an",
			required:       true,
			minLength:      3,
			maxLength:      6,
			expectedResult: false,
			expectedError:  BetweenLengthDefaultMessage,
		},
		{
			name:           "betweenTooLongAndRequired",
			value:          "avocado",
			required:       true,
			minLength:      3,
			maxLength:      6,
			expectedResult: false,
			expectedError:  BetweenLengthDefaultMessage,
		},
		{
			name:           "betweenTooLongAndNotRequired",
			value:          "avocado",
			required:       false,
			minLength:      3,
			maxLength:      6,
			expectedResult: false,
			expectedError:  BetweenLengthDefaultMessage,
		},
		{
			name:           "testCustomErrorMessage",
			value:          "avocado",
			required:       false,
			minLength:      3,
			maxLength:      6,
			expectedResult: false,
			expectedError:  "myCustomMessage",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := Between(tt.minLength, tt.maxLength, tt.expectedError)
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
