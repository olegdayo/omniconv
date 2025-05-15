package omniconv

// IntToBoolConverter converts int into bool
//
// Zero is converted to false, other values are converted to true
func IntToBoolConverter[T Int](from T) (to bool) {
	return from != 0
}

// BoolToIntConverter converts bool into int
func BoolToIntConverter[T Int](from bool) (to T) {
	if from {
		to++
	}
	return to
}
