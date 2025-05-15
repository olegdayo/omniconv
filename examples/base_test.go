package examples

import (
	"fmt"

	"github.com/olegdayo/omniconv"
)

func ExampleNumberConverter() {
	ints := []int{1, 2, 3, 4}
	floats := omniconv.ConvertSlice(ints, omniconv.NumberConverter[int, float64])
	fmt.Printf("%#v\n", floats)
	// Output: []float64{1, 2, 3, 4}
}

func ExampleIntToBoolConverter() {
	int8s := []int8{1, 0, 100, 0, -100, 0, 1, 0, -1}
	floats := omniconv.ConvertSlice(int8s, omniconv.IntToBoolConverter[int8])
	fmt.Printf("%#v\n", floats)
	// Output: []bool{true, false, true, false, true, false, true, false, true}
}

func ExampleStringToIntConverter() {
	strings := map[int]string{5: "6", 7: "8", 9: "silly"}
	uints := omniconv.ConvertMap(strings, omniconv.StringToIntConverter[int])
	fmt.Printf("%#v\n", uints)
	// Output:  map[int]int{5:6, 7:8, 9:0}
}

func ExampleStringToBoolConverter() {
	strings := map[int]string{5: "false", 7: "true", 9: "dummy"}
	uints := omniconv.ConvertMap(strings, omniconv.StringToBoolConverter)
	fmt.Printf("%#v\n", uints)
	// Output:  map[int]bool{5:false, 7:true, 9:false}
}
