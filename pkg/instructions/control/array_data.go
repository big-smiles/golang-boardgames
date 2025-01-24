package instruction_control

import "github.com/big-smiles/boardgame-golang/pkg/instruction"

type DataInstructionArray struct {
	dataInstructions []instruction.DataInstruction
}

func NewDataInstructionArray(d []instruction.DataInstruction) (*DataInstructionArray, error) {
	return &DataInstructionArray{
		dataInstructions: d,
	}, nil
}

func (d *DataInstructionArray) NewFromThisData() (instruction.Instruction, error) {
	i, err := newInstructionArray(*d)
	if err != nil {
		return nil, err
	}
	return i, nil
}
