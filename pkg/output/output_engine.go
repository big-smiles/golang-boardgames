package output

import (
	"github.com/big-smiles/golang-boardgames/pkg/entity"
	"github.com/big-smiles/golang-boardgames/pkg/player"
)

type Game struct {
	Entities             []entity.OutputEntity
	PropertyIds          entity.OutputManagerPropertyId
	CurrentActivePlayers []player.Id
}

func NewGameOutput(
	entities []entity.OutputEntity,
	propertyIds entity.OutputManagerPropertyId,
	currentActivePlayers []player.Id,
) (*Game, error) {
	return &Game{
		Entities:             entities,
		PropertyIds:          propertyIds,
		CurrentActivePlayers: currentActivePlayers,
	}, nil

}
