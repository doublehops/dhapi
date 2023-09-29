package validator

import (
	"reflect"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name           string
		rules          []Rule
		expectedErrors ErrorMessages
	}{
		{
			name: "allPass",
			rules: []Rule{
				{
					VariableName: "Make",
					Required:     true,
					Value:        "Ford",
					Function:     []ValidationFunctions{MinLength(4, "")},
				},
			},
			expectedErrors: ErrorMessages{},
		},
		{
			name: "makeFailMinLength",
			rules: []Rule{
				{
					VariableName: "Make",
					Required:     true,
					Value:        "Ford",
					Function:     []ValidationFunctions{MinLength(10, "")},
				},
			},
			expectedErrors: ErrorMessages{
				"Make": {MinLengthDefaultMessage},
			},
		},
		{
			name: "makePropertyEmptyButRequired",
			rules: []Rule{
				{
					VariableName: "Make",
					Required:     true,
					Value:        "",
					Function:     []ValidationFunctions{MinLength(10, "")},
				},
			},
			expectedErrors: ErrorMessages{
				"Make": {RequiredPropertyError},
			},
		},
		{
			name: "makeFailWithCustomError",
			rules: []Rule{
				{
					VariableName: "Make",
					Required:     true,
					Value:        "Ford",
					Function:     []ValidationFunctions{MinLength(10, "MyCustomerErrorMessage")},
				},
			},
			expectedErrors: ErrorMessages{
				"Make": {"MyCustomerErrorMessage"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			errors := RunValidation(tt.rules)
			if !reflect.DeepEqual(tt.expectedErrors, errors) {
				t.Errorf("Error not as expected. Expected: %v; got: %v", tt.expectedErrors, errors)
			}
		})
	}
}
