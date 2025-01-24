package interaction

import "github.com/big-smiles/golang-boardgames/pkg/instruction"

type ClearAvailableInteraction struct {
}

func (interaction *ClearAvailableInteraction) Execute(ctx instruction.ExecutionContext) error {
	err := ctx.Performer.Interaction.ClearAvailableInteraction()
	if err != nil {
		return err
	}
	return nil
}

func NewClearAvailableInteraction() (*ClearAvailableInteraction, error) {
	return &ClearAvailableInteraction{}, nil
}
