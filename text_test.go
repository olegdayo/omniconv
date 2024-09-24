package omniconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextConverterStringToBytes(t *testing.T) {
	type inputType = string
	type outputType = []byte

	testCases := []struct {
		name   string
		input  inputType
		output outputType
	}{
		{
			name:   "empty",
			input:  "",
			output: []byte{},
		},
		{
			name:   "non-empty",
			input:  "hello",
			output: []byte{'h', 'e', 'l', 'l', 'o'},
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
					TextConverter[inputType, outputType](testCase.input),
				)
			},
		)
	}
}
