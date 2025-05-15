# OmniConv

[ci]: https://github.com/olegdayo/omniconv/actions/workflows/ci.yaml/badge.svg

![CI][ci]

A simple, somewhat declarative type conversion library

# Examples

## Base Types

Include number, string, bool

```go
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
```

## Custom Types

```go
type ModelLogic struct {
	ID   int
	Name string
}

type ModelRepository struct {
	ID   int
	Name sql.NullString
}

func RepositoryToLogicConverter(from ModelRepository) (to ModelLogic) {
	to.ID = from.ID
	if from.Name.Valid {
		to.Name = from.Name.String
	}
	return to
}

func ExampleRepositoryToLogicConverter() {
	repositories := []ModelRepository{
		{
			ID: 123,
			Name: sql.NullString{
				String: "smth",
				Valid:  true,
			},
		},
		{
			ID:   456,
			Name: sql.NullString{},
		},
	}
	logics := omniconv.ConvertSlice(repositories, RepositoryToLogicConverter)
	fmt.Printf("%#v\n", logics)
	// Output: []examples.ModelLogic{examples.ModelLogic{ID:123, Name:"smth"}, examples.ModelLogic{ID:456, Name:""}}
}
```
