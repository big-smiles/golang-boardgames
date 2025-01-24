package output

import (
	"github.com/big-smiles/golang-boardgames/pkg/entity"
)

type Game struct {
	Entities []entity.OutputEntity
}

func NewGameOutput(
	entities []entity.OutputEntity,
) (*Game, error) {
	return &Game{
		Entities: entities,
	}, nil

}
