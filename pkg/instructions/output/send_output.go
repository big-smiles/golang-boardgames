package instructionOutput

import (
	"github.com/big-smiles/golang-boardgames/pkg/instruction"
)

type InstructionSendOutput struct {
}

func (i InstructionSendOutput) Execute(ctx instruction.ExecutionContext) error {
	err := ctx.Performer.Output.SendOutput()
	if err != nil {
		return err
	}
	return nil
}
func newInstructionSendOutput(_ DataInstructionSendOutput) (*InstructionSendOutput, error) {
	return &InstructionSendOutput{}, nil
}
