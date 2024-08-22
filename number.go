package converter

type SignedInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type UnsignedInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Int interface {
	SignedInt | UnsignedInt
}

type Float interface {
	~float32 | ~float64
}

type Number interface {
	Int | Float
}

func NumberConverter[T1 Number, T2 Number](from T1) (to T2) {
	return T2(from)
}
