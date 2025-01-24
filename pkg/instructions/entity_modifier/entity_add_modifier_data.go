package instructionEntityModifier

import (
	"github.com/big-smiles/boardgame-golang/pkg/entity"
	"github.com/big-smiles/boardgame-golang/pkg/instruction"
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

func NewDataInstructionAddEntityModifier(d entity.DataId, dataEntityModifier entity.DataModifier) (*DataInstructionAddEntityModifier, error) {

	return &DataInstructionAddEntityModifier{
		target:             d,
		dataEntityModifier: dataEntityModifier,
	}, nil
}
