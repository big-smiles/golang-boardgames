package instructionOutput

import "github.com/big-smiles/boardgame-golang/pkg/instruction"

type DataInstructionSendOutput struct{}

func NewDataInstructionSendOutput() (*DataInstructionSendOutput, error) {
	return &DataInstructionSendOutput{}, nil
}
func (d DataInstructionSendOutput) NewFromThisData() (instruction.Instruction, error) {
	i, er := newInstructionSendOutput(d)
	if er != nil {
		return nil, er
	}
	return i, nil
}
