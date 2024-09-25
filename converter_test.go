package omniconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertSlice(t *testing.T) {
	type inputType = int
	type outputType = string

	testCases := []struct {
		name   string
		input  []inputType
		output []outputType
	}{
		{
			name:   "empty",
			input:  []inputType{},
			output: []outputType{},
		},
		{
			name:   "non-empty",
			input:  []inputType{1, 2, 3},
			output: []outputType{"1", "2", "3"},
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
					ConvertSlice(testCase.input, IntToStringConverter[inputType]),
				)
			},
		)
	}
}

func TestConvertMap(t *testing.T) {
	type key = byte
	type inputType = int
	type outputType = string

	testCases := []struct {
		name   string
		input  map[key]inputType
		output map[key]outputType
	}{
		{
			name:   "empty",
			input:  map[key]inputType{},
			output: map[key]outputType{},
		},
		{
			name:   "non-empty",
			input:  map[key]inputType{16: 1, 32: 2, 64: 3},
			output: map[key]outputType{16: "1", 32: "2", 64: "3"},
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
					ConvertMap(testCase.input, IntToStringConverter[inputType]),
				)
			},
		)
	}
}
