package instruction_phase

import (
	"github.com/big-smiles/golang-boardgames/pkg/instruction"
	"github.com/big-smiles/golang-boardgames/pkg/phase"
)

type DataInstructionSetNextPhase struct {
	phaseName IValueResolver[phase.NamePhase]
}

func NewDataInstructionSetNextPhase(phaseName IValueResolver[phase.NamePhase]) *DataInstructionSetNextPhase {
	return &DataInstructionSetNextPhase{
		phaseName: phaseName,
	}
}

func (d *DataInstructionSetNextPhase) NewFromThisData() (instruction.Instruction, error) {
	i, err := newInstructionSetNextPhase(*d)
	if err != nil {
		return nil, err
	}
	return i, nil
}
