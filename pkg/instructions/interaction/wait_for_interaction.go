package interaction

import "github.com/big-smiles/boardgame-golang/pkg/instruction"

type WaitForInteraction struct {
}

func (interaction *WaitForInteraction) Execute(ctx instruction.ExecutionContext) error {
	err := ctx.Performer.Interaction.WaitForInteraction()
	if err != nil {
		return err
	}
	return nil
}

func NewWaitForInteraction() (*WaitForInteraction, error) {
	return &WaitForInteraction{}, nil
}
