package converter

type Text interface {
	[]byte | string
}

// NumberConverter converts text type T1 to text type T2
func TextConverter[T1 Text, T2 Text](from T1) (to T2) {
	return T2(from)
}
