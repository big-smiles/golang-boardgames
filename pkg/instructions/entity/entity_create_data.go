package instruction_entity

import (
	entity "github.com/big-smiles/boardgame-golang/pkg/entity"
	"github.com/big-smiles/boardgame-golang/pkg/instruction"
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
