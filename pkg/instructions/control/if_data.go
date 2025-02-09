package instruction_control

import "github.com/big-smiles/golang-boardgames/pkg/instruction"

type DataInstructionIf struct {
	condition IValueResolver[bool]
	whenTrue  instruction.DataInstruction
	whenFalse instruction.DataInstruction
}

func NewDataInstructionIf(
	condition IValueResolver[bool],
	whenTrue instruction.DataInstruction,
	whenFalse instruction.DataInstruction,
) *DataInstructionIf {
	return &DataInstructionIf{
		condition: condition,
		whenTrue:  whenTrue,
		whenFalse: whenFalse,
	}
}

func (d *DataInstructionIf) NewFromThisData() (instruction.Instruction, error) {
	i, err := newInstructionIf(*d)
	if err != nil {
		return nil, err
	}
	return i, nil
}
