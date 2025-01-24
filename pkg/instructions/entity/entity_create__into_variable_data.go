package instruction_entity

import (
	"github.com/big-smiles/boardgame-golang/pkg/entity"
	"github.com/big-smiles/boardgame-golang/pkg/instruction"
)

type DataInstructionCreateEntityIntoVariable struct {
	dataEntity           entity.NameDataEntity
	variablePropertyName entity.NamePropertyId[entity.Id]
}

func (d DataInstructionCreateEntityIntoVariable) NewFromThisData() (instruction.Instruction, error) {
	i, err := newInstructionCreateEntityIntoVariable(d)
	if err != nil {
		return nil, err
	}
	return i, nil
}

func NewDataInstructionCreateEntityIntoVariable(
	d entity.NameDataEntity,
	variablePropertyName entity.NamePropertyId[entity.Id],
) (*DataInstructionCreateEntityIntoVariable, error) {

	return &DataInstructionCreateEntityIntoVariable{
		dataEntity:           d,
		variablePropertyName: variablePropertyName,
	}, nil
}
