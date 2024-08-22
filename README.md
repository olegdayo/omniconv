# Converter

[ci]: https://github.com/olegdayo/converter/actions/workflows/ci.yaml/badge.svg

![CI][ci]

A simple, somewhat declarative type conversion library

# Examples

## Base Types

```go
func BaseExamples() {
	ints := []int{1, 2, 3, 4}
	floats := converter.ConvertSlice(ints, converter.NumberConverter[int, float64])
	_ = floats // []float64{1, 2, 3, 4}

	strings := map[int]string{5: "6", 7: "8", 9: "silly"}
	uints := converter.ConvertMap(strings, converter.StringToIntConverter[uint])
	_ = uints // map[int]uint{5: 6, 7: 8, 9: 0}
}
```

## Custom Types

```go
type ModelLogic struct {
	ID   int
	Name *string
}

type ModelRepository struct {
	ID   int
	Name sql.NullString
}

func RepositoryToLogicConverter(from ModelRepository) (to ModelLogic) {
	to.ID = from.ID
	if from.Name.Valid {
		to.Name = &from.Name.String
	}
	return to
}

func CustomExamples() {
	repositories := []ModelRepository{
		{
			ID: 0xD00D,
			Name: sql.NullString{
				String: "smth",
				Valid:  true,
			},
		},
		{
			ID:   0xCA2E,
			Name: sql.NullString{},
		},
	}
	logics := converter.ConvertSlice(repositories, RepositoryToLogicConverter)
	_ = logics // []ModelLogic{{ID: 0xD00D, Name: &"smth"}, {ID: 0xCA2E, Name: nil}}
}
```
