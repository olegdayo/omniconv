package omniconv

// ConvertSlice converts slice of type F to slice of type T
func ConvertSlice[F any, T any](from []F, f func(F) T) (to []T) {
	to = make([]T, len(from))
	for i, v := range from {
		to[i] = f(v)
	}
	return to
}

// ConvertMap converts map keys FK and values FV to map keys TK and values TV
func ConvertMap[FK comparable, TK comparable, FV any, TV any](from map[FK]FV, fk func(FK) TK, fv func(FV) TV) (to map[TK]TV) {
	to = make(map[TK]TV, len(from))
	for i, v := range from {
		to[fk(i)] = fv(v)
	}
	return to
}

// ConvertMapValues converts map values of type F to map values of type T
func ConvertMapValues[K comparable, F any, T any](from map[K]F, f func(F) T) (to map[K]T) {
	to = make(map[K]T, len(from))
	for i, v := range from {
		to[i] = f(v)
	}
	return to
}

// ConvertMapKeys converts map keys of type F to map keys of type T
func ConvertMapKeys[F comparable, T comparable, V any](from map[F]V, f func(F) T) (to map[T]V) {
	to = make(map[T]V, len(from))
	for i, v := range from {
		to[f(i)] = v
	}
	return to
}
