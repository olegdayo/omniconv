package omniconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToStringConverter(t *testing.T) {
	type inputType = int
	type outputType = string

	testCases := []struct {
		name   string
		input  inputType
		output outputType
	}{
		{
			name:   "empty",
			input:  0,
			output: "0",
		},
		{
			name:   "non-empty",
			input:  1,
			output: "1",
		},
		{
			name:   "non-empty",
			input:  -1,
			output: "-1",
		},
		{
			name:   "non-empty",
			input:  0b1111,
			output: "15",
		},
		{
			name:   "non-empty",
			input:  0o100,
			output: "64",
		},
		{
			name:   "non-empty",
			input:  0xFF,
			output: "255",
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
					IntToStringConverter(testCase.input),
				)
			},
		)
	}
}

func TestStringToIntConverter(t *testing.T) {
	type inputType = string
	type outputType = int

	testCases := []struct {
		name   string
		input  inputType
		output outputType
	}{
		{
			name:   "empty",
			input:  "0",
			output: 0,
		},
		{
			name:   "non-empty",
			input:  "1",
			output: 1,
		},
		{
			name:   "non-empty",
			input:  "-1",
			output: -1,
		},
		{
			name:   "non-empty",
			input:  "0xFF",
			output: 0,
		},
		{
			name:   "non-empty",
			input:  "1.5",
			output: 0,
		},
		{
			name:   "non-empty",
			input:  "LOL",
			output: 0,
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
					StringToIntConverter[outputType](testCase.input),
				)
			},
		)
	}
}
