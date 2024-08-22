package converter

type Text interface {
	[]byte | string
}

func TextConverter[T1 Text, T2 Text](from T1) (to T2) {
	return T2(from)
}
