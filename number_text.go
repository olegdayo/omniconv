package converter

import (
	"fmt"
	"strconv"
)

func IntToStringConverter[T Int](from T) (to string) {
	return fmt.Sprintf("%d", from)
}

func FloatToStringConverter[T Float](from T) (to string) {
	return fmt.Sprintf("%f", from)
}

func StringToIntConverter[T Int](from string) (to T) {
	n, err := strconv.Atoi(from)
	if err != nil {
		return 0
	}
	return T(n)
}

func StringToFloatConverter[T Float](from string) (to T) {
	const bitSize = 64
	n, err := strconv.ParseFloat(from, bitSize)
	if err != nil {
		return 0
	}
	return T(n)
}
