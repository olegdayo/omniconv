package omniconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	epsilon = 0.0000000001
)

func isFloatEqual[T Float](f T, s T) bool {
	diff := f - s
	if diff < 0 {
		diff = s - f
	}
	return diff < epsilon
}

func TestNumberConverterIntToFloat(t *testing.T) {
	type inputType = float64
	type outputType = int

	testCases := []struct {
		name   string
		input  inputType
		output outputType
	}{
		{
			name:   "empty",
			input:  0,
			output: 0,
		},
		{
			name:   "non-empty",
			input:  1,
			output: 1,
		},
		{
			name:   "non-empty",
			input:  -1,
			output: -1,
		},
		{
			name:   "non-empty",
			input:  1.5,
			output: 1,
		},
		{
			name:   "non-empty",
			input:  -1.5,
			output: -1,
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
					NumberConverter[inputType, outputType](testCase.input),
				)
			},
		)
	}
}

func TestNumberConverterFloatToInt(t *testing.T) {
	type inputType = int
	type outputType = float64

	testCases := []struct {
		name   string
		input  inputType
		output outputType
	}{
		{
			name:   "empty",
			input:  0,
			output: 0,
		},
		{
			name:   "non-empty",
			input:  1,
			output: 1,
		},
		{
			name:   "non-empty",
			input:  -1,
			output: -1,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(
			testCase.name,
			func(t *testing.T) {
				t.Parallel()
				assert.True(
					t,
					isFloatEqual(testCase.output, NumberConverter[inputType, outputType](testCase.input)),
				)
			},
		)
	}
}

func TestNumberConverterIntToUint(t *testing.T) {
	type inputType = int
	type outputType = uint

	testCases := []struct {
		name   string
		input  inputType
		output outputType
	}{
		{
			name:   "empty",
			input:  0,
			output: 0,
		},
		{
			name:   "non-empty",
			input:  1,
			output: 1,
		},
		{
			name:   "non-empty",
			input:  -1,
			output: 0xffffffffffffffff, // ATTENTION!
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
					NumberConverter[inputType, outputType](testCase.input),
				)
			},
		)
	}
}
