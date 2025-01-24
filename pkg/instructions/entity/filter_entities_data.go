package instruction_entity

import (
	"github.com/big-smiles/boardgame-golang/pkg/entity"
	"github.com/big-smiles/boardgame-golang/pkg/instruction"
)

type DataInstructionFilterEntities struct {
	predicate      entity.Predicate
	namePropertyId entity.NamePropertyId[[]entity.Id]
}

func NewDataInstructionFilterEntities(
	predicate entity.Predicate,
	namePropertyId entity.NamePropertyId[[]entity.Id],
) (*DataInstructionFilterEntities, error) {
	return &DataInstructionFilterEntities{
		predicate:      predicate,
		namePropertyId: namePropertyId,
	}, nil
}
func (d DataInstructionFilterEntities) NewFromThisData() (instruction.Instruction, error) {
	return NewInstructionFilterEntities(d)
}
