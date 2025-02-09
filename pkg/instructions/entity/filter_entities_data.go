package instruction_entity

import (
	"github.com/big-smiles/golang-boardgames/pkg/entity"
	"github.com/big-smiles/golang-boardgames/pkg/instruction"
)

type DataInstructionFilterEntities struct {
	predicate      entity.Predicate
	namePropertyId entity.NamePropertyId[[]entity.Id]
}

func NewDataInstructionFilterEntities(
	predicate entity.Predicate,
	namePropertyId entity.NamePropertyId[[]entity.Id],
) *DataInstructionFilterEntities {
	return &DataInstructionFilterEntities{
		predicate:      predicate,
		namePropertyId: namePropertyId,
	}
}
func (d DataInstructionFilterEntities) NewFromThisData() (instruction.Instruction, error) {
	return NewInstructionFilterEntities(d)
}
