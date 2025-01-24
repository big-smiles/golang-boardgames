package interaction

import (
	"github.com/big-smiles/golang-boardgames/pkg/entity"
	"github.com/big-smiles/golang-boardgames/pkg/player"
)

type OutputInteraction struct {
	Id                Id
	PlayerId          player.Id
	AvailableEntities []entity.Id
	//MinAmount minimum amount of entities that can be selected INCLUSIVE
	MinAmount int
	//MaxAmount maximum amount of entities that can be selected EXCLUSIVE
	MaxAmount int
}

func NewOutputInteraction(data AvailableInteraction) (*OutputInteraction, error) {
	return &OutputInteraction{
		Id:                data.Id,
		PlayerId:          data.PlayerId,
		AvailableEntities: data.AvailableEntities,
		MinAmount:         data.MinAmount,
		MaxAmount:         data.MaxAmount,
	}, nil
}
