# Converter

[ci]: https://github.com/olegdayo/converter/actions/workflows/ci.yaml/badge.svg

![CI][ci]

A simple, somewhat declarative type conversion library

# Examples

## Base Types

```go
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

func ExampleCustom() {
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
	logics := converter.ConvertSlice(repositories, RepositoryToLogicConverter)
	fmt.Printf("%#v\n", logics)
	// Output: []examples.ModelLogic{examples.ModelLogic{ID:123, Name:"smth"}, examples.ModelLogic{ID:456, Name:""}}
}
```
