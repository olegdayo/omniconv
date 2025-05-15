package omniconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToBoolConverter(t *testing.T) {
	type inputType = int
	type outputType = bool

	testCases := []struct {
		name   string
		input  inputType
		output outputType
	}{
		{
			name:   "empty",
			input:  0,
			output: false,
		},
		{
			name:   "positive",
			input:  1,
			output: true,
		},
		{
			name:   "other-positive",
			input:  255,
			output: true,
		},
		{
			name:   "negative",
			input:  -1,
			output: true,
		},
		{
			name:   "other-negative",
			input:  -256,
			output: true,
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
					IntToBoolConverter(testCase.input),
				)
			},
		)
	}
}

func TestBoolToIntConverter(t *testing.T) {
	type inputType = bool
	type outputType = int

	testCases := []struct {
		name   string
		input  inputType
		output outputType
	}{
		{
			name:   "false",
			input:  false,
			output: 0,
		},
		{
			name:   "true",
			input:  true,
			output: 1,
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
					BoolToIntConverter[outputType](testCase.input),
				)
			},
		)
	}
}
