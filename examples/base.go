package examples

import (
	"github.com/olegdayo/converter"
)

func BaseExamples() {
	ints := []int{1, 2, 3, 4}
	floats := converter.ConvertSlice(ints, converter.NumberConverter[int, float64])
	_ = floats // []float64{1, 2, 3, 4}

	strings := map[int]string{5: "6", 7: "8", 9: "silly"}
	uints := converter.ConvertMap(strings, converter.StringToIntConverter[uint])
	_ = uints // map[int]uint{5: 6, 7: 8, 9: 0}
}
