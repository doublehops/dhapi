package validator

import (
	"reflect"
	"testing"

	"github.com/doublehops/dhapi/responses"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name           string
		rules          []Rule
		expectedErrors responses.ErrorMessages
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
			expectedErrors: responses.ErrorMessages{},
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
			expectedErrors: responses.ErrorMessages{
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
			expectedErrors: responses.ErrorMessages{
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
			expectedErrors: responses.ErrorMessages{
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
