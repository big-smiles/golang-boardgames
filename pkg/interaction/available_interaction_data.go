package interaction

import (
	"github.com/big-smiles/golang-boardgames/pkg/entity"
	"github.com/big-smiles/golang-boardgames/pkg/player"
)

type DataAvailableInteraction struct {
	playerId          player.Id
	availableEntities IValueResolver[[]entity.Id]
	//MinAmount minimum amount of entities that can be selected INCLUSIVE
	minAmount int
	//MaxAmount maximum amount of entities that can be selected EXCLUSIVE
	maxAmount int
}

// NewDataAvailableInteraction
// minAmount is an INCLUSIVE value, maxAmount is an EXCLUSIVE value
func NewDataAvailableInteraction(
	playerId player.Id,
	availableEntities IValueResolver[[]entity.Id],
	minAmount int,
	maxAmount int,
) (*DataAvailableInteraction, error) {
	return &DataAvailableInteraction{
		playerId:          playerId,
		availableEntities: availableEntities,
		minAmount:         minAmount,
		maxAmount:         maxAmount,
	}, nil
}
