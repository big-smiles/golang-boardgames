package instructionOutput

import "github.com/big-smiles/golang-boardgames/pkg/instruction"

type DataInstructionSendOutput struct{}

func NewDataInstructionSendOutput() *DataInstructionSendOutput {
	return &DataInstructionSendOutput{}
}
func (d DataInstructionSendOutput) NewFromThisData() (instruction.Instruction, error) {
	i, er := newInstructionSendOutput(d)
	if er != nil {
		return nil, er
	}
	return i, nil
}
