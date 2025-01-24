package instruction_entity

import (
	"github.com/big-smiles/boardgame-golang/pkg/entity"
	"github.com/big-smiles/boardgame-golang/pkg/instruction"
)

type InstructionFilterEntities struct {
	predicate      entity.Predicate
	namePropertyId entity.NamePropertyId[[]entity.Id]
}

func NewInstructionFilterEntities(data DataInstructionFilterEntities) (*InstructionFilterEntities, error) {
	return &InstructionFilterEntities{
		predicate:      data.predicate,
		namePropertyId: data.namePropertyId,
	}, nil
}

func (i InstructionFilterEntities) Execute(ctx instruction.ExecutionContext) error {
	err := ctx.Performer.Entity.FilterEntitiesIntoVariable(ctx.ExecutionVariables, i.predicate, i.namePropertyId)
	if err != nil {
		return err
	}
	return nil
}
