package omniconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoolToStringConverter(t *testing.T) {
	type inputType = bool
	type outputType = string

	testCases := []struct {
		name   string
		input  inputType
		output outputType
	}{
		{
			name:   "false",
			input:  false,
			output: "false",
		},
		{
			name:   "true",
			input:  true,
			output: "true",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(
			testCase.name,
			func(t *testing.T) {
				t.Parallel()
				assert.Equal(
					t,
					testCase.output,
					BoolToStringConverter(testCase.input),
				)
			},
		)
	}
}

func TestStringToBoolConverter(t *testing.T) {
	type inputType = string
	type outputType = bool

	testCases := []struct {
		name   string
		input  inputType
		output outputType
	}{
		{
			name:   "true",
			input:  "true",
			output: true,
		},
		{
			name:   "false",
			input:  "false",
			output: false,
		},
		{
			name:   "wrong",
			input:  "lol",
			output: false,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(
			testCase.name,
			func(t *testing.T) {
				t.Parallel()
				assert.Equal(
					t,
					testCase.output,
					StringToBoolConverter(testCase.input),
				)
			},
		)
	}
}
