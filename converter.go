package converter

func ConvertSlice[F any, T any](from []F, f func(F) T) (to []T) {
	to = make([]T, len(from))
	for i, v := range from {
		to[i] = f(v)
	}
	return to
}

func ConvertMap[K comparable, F any, T any](from map[K]F, f func(F) T) (to map[K]T) {
	to = make(map[K]T, len(from))
	for i, v := range from {
		to[i] = f(v)
	}
	return to
}
