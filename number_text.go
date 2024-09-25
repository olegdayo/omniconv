package omniconv

import (
	"fmt"
	"strconv"
)

// IntToStringConverter converts integer into string
func IntToStringConverter[T Int](from T) (to string) {
	return fmt.Sprintf("%d", from)
}

// FloatToStringConverter converts float into string
func FloatToStringConverter[T Float](from T) (to string) {
	return fmt.Sprintf("%f", from)
}

// StringToIntConverter parses int from string
//
// On error returns default value (0)
func StringToIntConverter[T Int](from string) (to T) {
	n, err := strconv.Atoi(from)
	if err != nil {
		return 0
	}
	return T(n)
}

// MustStringToIntConverter parses int from string
//
// On error panics
func MustStringToIntConverter[T Int](from string) (to T) {
	n, err := strconv.Atoi(from)
	if err != nil {
		panic(err)
	}
	return T(n)
}

// StringToFloatConverter parses float from string
//
// On error returns default value (0)
func StringToFloatConverter[T Float](from string) (to T) {
	const bitSize = 64
	n, err := strconv.ParseFloat(from, bitSize)
	if err != nil {
		return 0
	}
	return T(n)
}

// MustStringToFloatConverter parses float from string
//
// On error panics
func MustStringToFloatConverter[T Float](from string) (to T) {
	const bitSize = 64
	n, err := strconv.ParseFloat(from, bitSize)
	if err != nil {
		panic(err)
	}
	return T(n)
}
