package instructionEntityModifier

import (
	"github.com/big-smiles/golang-boardgames/pkg/entity"
	"github.com/big-smiles/golang-boardgames/pkg/instruction"
)

type DataInstructionAddEntityModifier struct {
	target             entity.DataId
	dataEntityModifier entity.DataModifier
}

func (d DataInstructionAddEntityModifier) NewFromThisData() (instruction.Instruction, error) {
	i, err := newInstructionAddEntityModifier(d)
	if err != nil {
		return nil, err
	}
	return i, nil
}

func NewDataInstructionAddEntityModifier(
	d entity.DataId,
	dataEntityModifier entity.DataModifier,
) *DataInstructionAddEntityModifier {
	return &DataInstructionAddEntityModifier{
		target:             d,
		dataEntityModifier: dataEntityModifier,
	}
}
