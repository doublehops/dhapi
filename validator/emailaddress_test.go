package validator

import "testing"

func TestEmailAddress(t *testing.T) {
	tests := []struct {
		name           string
		value          string
		required       bool
		expectedResult bool
		expectedError  string
	}{
		{
			name:           "validEmailAddress",
			value:          "test@example.com",
			required:       true,
			expectedResult: true,
			expectedError:  "",
		},
		{
			name:           "invalidEmailAddressButNotRequired",
			value:          "test.last",
			required:       false,
			expectedResult: false,
			expectedError:  EmailAddressDefaultMessage,
		},
		{
			name:           "emptyEmailAddressButIsRequired",
			value:          "",
			required:       true,
			expectedResult: false,
			expectedError:  RequiredPropertyError,
		},
		{
			name:           "emptyEmailAddressButNotRequired",
			value:          "",
			required:       false,
			expectedResult: true,
			expectedError:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := EmailAddress(tt.expectedError)
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
