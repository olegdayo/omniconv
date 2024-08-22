package examples

import (
	"database/sql"
	"fmt"

	"github.com/olegdayo/converter"
)

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
