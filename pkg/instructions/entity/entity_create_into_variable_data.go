package instruction_entity

import (
	"github.com/big-smiles/golang-boardgames/pkg/entity"
	"github.com/big-smiles/golang-boardgames/pkg/instruction"
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
) *DataInstructionCreateEntityIntoVariable {

	return &DataInstructionCreateEntityIntoVariable{
		dataEntity:           d,
		variablePropertyName: variablePropertyName,
	}
}
