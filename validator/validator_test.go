package validator

import (
	"reflect"
	"testing"

	"github.com/doublehops/dhapi/resp"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name           string
		rules          []Rule
		expectedErrors resp.ErrMsgs
	}{
		{
			name: "allPass",
			rules: []Rule{
				{
					VariableName: "Make",
					Required:     true,
					Value:        "Ford",
					Function:     []ValidateFuncs{MinLength(4, "")},
				},
			},
			expectedErrors: resp.ErrMsgs{},
		},
		{
			name: "makeFailMinLength",
			rules: []Rule{
				{
					VariableName: "Make",
					Required:     true,
					Value:        "Ford",
					Function:     []ValidateFuncs{MinLength(10, "")},
				},
			},
			expectedErrors: resp.ErrMsgs{
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
					Function:     []ValidateFuncs{MinLength(10, "")},
				},
			},
			expectedErrors: resp.ErrMsgs{
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
					Function:     []ValidateFuncs{MinLength(10, "MyCustomerErrorMessage")},
				},
			},
			expectedErrors: resp.ErrMsgs{
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
