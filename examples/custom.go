package examples

import (
	"database/sql"

	"github.com/olegdayo/converter"
)

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
