package instruction_control

import "github.com/big-smiles/golang-boardgames/pkg/instruction"

type DataInstructionArray struct {
	dataInstructions []instruction.DataInstruction
}

func NewDataInstructionArray(d ...instruction.DataInstruction) *DataInstructionArray {
	return &DataInstructionArray{
		dataInstructions: d,
	}
}

func (d *DataInstructionArray) NewFromThisData() (instruction.Instruction, error) {
	i, err := newInstructionArray(*d)
	if err != nil {
		return nil, err
	}
	return i, nil
}
