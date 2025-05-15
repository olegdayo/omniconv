package omniconv

import (
	"fmt"
	"strconv"
)

// BoolToStringConverter converts bool into string
func BoolToStringConverter(from bool) (to string) {
	return fmt.Sprintf("%t", from)
}

// StringToBoolConverter parses bool from string
//
// On error returns default value (false)
func StringToBoolConverter(from string) (to bool) {
	to, err := strconv.ParseBool(from)
	if err != nil {
		return false
	}
	return to
}

// MustStringToBoolConverter parses bool from string
//
// On error panics
func MustStringToBoolConverter(from string) (to bool) {
	to, err := strconv.ParseBool(from)
	if err != nil {
		panic(err)
	}
	return to
}
