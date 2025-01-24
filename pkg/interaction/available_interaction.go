package interaction

import (
	"github.com/big-smiles/boardgame-golang/pkg/entity"
	"github.com/big-smiles/boardgame-golang/pkg/player"
)

type AvailableInteraction struct {
	Id                Id
	PlayerId          player.Id
	AvailableEntities []entity.Id
	//MinAmount minimum amount of entities that can be selected INCLUSIVE
	MinAmount int
	//MaxAmount maximum amount of entities that can be selected EXCLUSIVE
	MaxAmount int
	//TODO:this can be turned into a typed ID in this package, the Game package should be the one
	// calling the trigger on the instruction package with the interaction Id
	instructionIdToTrigger int
}

func NewAvailableInteraction(
	executionVariables entity.Entity,
	managerPropertyId *entity.ManagerPropertyId,
	id Id,
	instructionIdToTrigger int,
	data DataAvailableInteraction,
) (*AvailableInteraction, error) {
	availableEntities, err := data.availableEntities.Resolve(executionVariables, managerPropertyId)
	if err != nil {
		return nil, err
	}

	return &AvailableInteraction{
		Id:                     id,
		PlayerId:               data.playerId,
		AvailableEntities:      availableEntities,
		MinAmount:              data.minAmount,
		MaxAmount:              data.maxAmount,
		instructionIdToTrigger: instructionIdToTrigger,
	}, nil
}
