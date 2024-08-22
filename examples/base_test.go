package examples

import (
	"fmt"

	"github.com/olegdayo/converter"
)

func ExampleNumber() {
	ints := []int{1, 2, 3, 4}
	floats := converter.ConvertSlice(ints, converter.NumberConverter[int, float64])
	fmt.Printf("%#v\n", floats)
	// Output: []float64{1, 2, 3, 4}
}

func ExampleString() {
	strings := map[int]string{5: "6", 7: "8", 9: "silly"}
	uints := converter.ConvertMap(strings, converter.StringToIntConverter[int])
	fmt.Printf("%#v\n", uints)
	// Output:  map[int]int{5:6, 7:8, 9:0}
}
