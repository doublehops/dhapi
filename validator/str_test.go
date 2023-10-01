package validator

import "testing"

type testType string

func TestIn(t *testing.T) {
	tests := []struct {
		name           string
		value          string
		required       bool
		testSlice      []interface{}
		expectedResult bool
		expectedError  string
	}{
		{
			name:           "valueFoundInSlice",
			value:          "banana",
			required:       true,
			testSlice:      []interface{}{"apple", "banana", "carrot"},
			expectedResult: true,
			expectedError:  "",
		},
		{
			name:           "valueNotFoundInSlice",
			value:          "watermelon",
			required:       true,
			testSlice:      []interface{}{"apple", "banana", "carrot"},
			expectedResult: false,
			expectedError:  keyNotInSlice,
		},
		{
			name:           "valueNotFoundInSliceButNotRequired",
			value:          "watermelon",
			required:       false,
			testSlice:      []interface{}{"apple", "banana", "carrot"},
			expectedResult: false,
			expectedError:  keyNotInSlice,
		},
		{
			name:           "valueBlankAndRequired",
			value:          "",
			required:       true,
			testSlice:      []interface{}{"apple", "banana", "carrot"},
			expectedResult: false,
			expectedError:  RequiredPropertyError,
		},
		{
			name:           "valueBlankButNotRequired",
			value:          "",
			required:       false,
			testSlice:      []interface{}{"apple", "banana", "carrot"},
			expectedResult: true,
			expectedError:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := In(tt.testSlice, "")
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
