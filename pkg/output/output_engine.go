package output

import (
	"github.com/big-smiles/golang-boardgames/pkg/entity"
)

type Game struct {
	Entities    []entity.OutputEntity
	PropertyIds entity.OutputManagerPropertyId
}

func NewGameOutput(
	entities []entity.OutputEntity,
	propertyIds entity.OutputManagerPropertyId,
) (*Game, error) {
	return &Game{
		Entities:    entities,
		PropertyIds: propertyIds,
	}, nil

}
