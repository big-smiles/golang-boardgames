package instruction_control

import (
	"github.com/big-smiles/golang-boardgames/pkg/instruction"
)

type InstructionArray struct {
	instructions []instruction.Instruction
}

func (i InstructionArray) Execute(ctx instruction.ExecutionContext) error {
	for _, v := range i.instructions {
		err := v.Execute(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func newInstructionArray(d DataInstructionArray) (*InstructionArray, error) {
	instructions := make([]instruction.Instruction, len(d.dataInstructions))
	for i, v := range d.dataInstructions {
		r, err := v.NewFromThisData()
		if err != nil {
			return nil, err
		}
		instructions[i] = r
	}
	return &InstructionArray{
		instructions: instructions,
	}, nil
}
