package validator

import (
	"testing"
)

func TestMinValue(t *testing.T) {
	tests := []struct {
		name           string
		value          int
		required       bool
		minValue       int
		expectedResult bool
		expectedError  string
	}{
		{
			name:           "minValuePass",
			value:          11,
			required:       true,
			minValue:       10,
			expectedResult: true,
			expectedError:  "",
		},
		{
			name:           "minValuePassEqual",
			value:          10,
			required:       true,
			minValue:       10,
			expectedResult: true,
			expectedError:  "",
		},
		{
			name:           "minValueFail",
			value:          1,
			required:       true,
			minValue:       10,
			expectedResult: false,
			expectedError:  MinValueDefaultMessage,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := MinValue(tt.minValue, tt.expectedError)
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

func TestMaxValue(t *testing.T) {
	tests := []struct {
		name           string
		value          int
		required       bool
		maxValue       int
		expectedResult bool
		expectedError  string
	}{
		{
			name:           "maxValuePass",
			value:          1,
			required:       true,
			maxValue:       10,
			expectedResult: true,
			expectedError:  "",
		},
		{
			name:           "maxValuePassEqual",
			value:          10,
			required:       true,
			maxValue:       10,
			expectedResult: true,
			expectedError:  "",
		},
		{
			name:           "maxValueFail",
			value:          11,
			required:       true,
			maxValue:       10,
			expectedResult: false,
			expectedError:  MaxValueDefaultMessage,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := MaxValue(tt.maxValue, tt.expectedError)
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

func TestInInRange(t *testing.T) {
	tests := []struct {
		name           string
		value          int
		required       bool
		minValue       int
		maxValue       int
		expectedResult bool
		expectedError  string
	}{
		{
			name:           "inRangeLow",
			value:          3,
			required:       true,
			minValue:       3,
			maxValue:       6,
			expectedResult: true,
			expectedError:  "",
		},
		{
			name:           "inRangeHigh",
			value:          6,
			required:       false,
			minValue:       3,
			maxValue:       6,
			expectedResult: true,
			expectedError:  "",
		},
		{
			name:           "belowRangeFail",
			value:          2,
			required:       false,
			minValue:       3,
			maxValue:       6,
			expectedResult: false,
			expectedError:  InRangeDefaultMessage,
		},
		{
			name:           "aboveRangeFail",
			value:          7,
			required:       false,
			minValue:       3,
			maxValue:       6,
			expectedResult: false,
			expectedError:  InRangeDefaultMessage,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := IntInRange(tt.minValue, tt.maxValue, tt.expectedError)
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
