package instruction_entity

import (
	entity "github.com/big-smiles/golang-boardgames/pkg/entity"
	"github.com/big-smiles/golang-boardgames/pkg/instruction"
)

type DataInstructionCreateEntity struct {
	dataEntity entity.NameDataEntity
}

func (d DataInstructionCreateEntity) NewFromThisData() (instruction.Instruction, error) {
	i, err := newInstructionCreateEntity(d)
	if err != nil {
		return nil, err
	}
	return i, nil
}

func NewDataInstructionCreateEntity(d entity.NameDataEntity) (*DataInstructionCreateEntity, error) {

	return &DataInstructionCreateEntity{
		dataEntity: d,
	}, nil
}
