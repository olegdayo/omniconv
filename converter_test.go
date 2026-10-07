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
	type inputTypeKey = int
	type outputTypeKey = string
	type inputTypeValue = bool
	type outputTypeValue = int

	testCases := []struct {
		name   string
		input  map[inputTypeKey]inputTypeValue
		output map[outputTypeKey]outputTypeValue
	}{
		{
			name:   "empty",
			input:  map[inputTypeKey]inputTypeValue{},
			output: map[outputTypeKey]outputTypeValue{},
		},
		{
			name:   "non-empty",
			input:  map[inputTypeKey]inputTypeValue{16: true, 32: false, 64: true},
			output: map[outputTypeKey]outputTypeValue{"16": 1, "32": 0, "64": 1},
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
					ConvertMap(testCase.input, IntToStringConverter[inputTypeKey], BoolToIntConverter[outputTypeValue]),
				)
			},
		)
	}
}

func TestConvertMapValues(t *testing.T) {
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
					ConvertMapValues(testCase.input, IntToStringConverter[inputType]),
				)
			},
		)
	}
}

func TestConvertMapKeys(t *testing.T) {
	type value = byte
	type inputType = int
	type outputType = string

	testCases := []struct {
		name   string
		input  map[inputType]value
		output map[outputType]value
	}{
		{
			name:   "empty",
			input:  map[inputType]value{},
			output: map[outputType]value{},
		},
		{
			name:   "non-empty",
			input:  map[inputType]value{16: 1, 32: 2, 64: 3},
			output: map[outputType]value{"16": 1, "32": 2, "64": 3},
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
					ConvertMapKeys(testCase.input, IntToStringConverter[inputType]),
				)
			},
		)
	}
}
