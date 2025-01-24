package interaction

import (
	"github.com/big-smiles/golang-boardgames/pkg/entity"
	"github.com/big-smiles/golang-boardgames/pkg/player"
)

// SelectedInteraction an interaction taken by a PlayerId
type SelectedInteraction struct {
	id               Id
	playerId         player.Id
	selectedEntities []entity.Id
}

// NewSelectedInteraction when a PlayerId chooses an interaction and its targets create a
// SelectedInteraction to notify the engine
func NewSelectedInteraction(
	id Id,
	playerId player.Id,
	selectedEntities []entity.Id,
) (*SelectedInteraction, error) {
	return &SelectedInteraction{
		id:               id,
		playerId:         playerId,
		selectedEntities: selectedEntities,
	}, nil
}
