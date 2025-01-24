package output

import (
	"github.com/big-smiles/boardgame-golang/pkg/entity"
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
