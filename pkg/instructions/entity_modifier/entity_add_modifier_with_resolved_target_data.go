package instructionEntityModifier

import (
	"github.com/big-smiles/boardgame-golang/pkg/entity"
	"github.com/big-smiles/boardgame-golang/pkg/instruction"
)

type DataInstructionAddEntityModifierWithResolvedTarget struct {
	target             IValueResolver[[]entity.Id]
	dataEntityModifier entity.DataModifier
}

func (d DataInstructionAddEntityModifierWithResolvedTarget) NewFromThisData() (instruction.Instruction, error) {
	i, err := newInstructionAddEntityModifierWithResolvedTarget(d)
	if err != nil {
		return nil, err
	}
	return i, nil
}

func NewDataInstructionAddEntityModifierWithResolvedTarget(
	d IValueResolver[[]entity.Id],
	dataEntityModifier entity.DataModifier,
) (*DataInstructionAddEntityModifierWithResolvedTarget, error) {

	return &DataInstructionAddEntityModifierWithResolvedTarget{
		target:             d,
		dataEntityModifier: dataEntityModifier,
	}, nil
}
