package interaction

import (
	"github.com/big-smiles/boardgame-golang/pkg/instruction"
	"github.com/big-smiles/boardgame-golang/pkg/interaction"
)

type AddAvailableInteraction struct {
	availableInteraction interaction.DataAvailableInteraction
	dataInstruction      instruction.DataInstruction
}

func (a AddAvailableInteraction) Execute(ctx instruction.ExecutionContext) error {
	err := ctx.Performer.Interaction.AddAvailableInteraction(
		ctx.ExecutionVariables,
		a.availableInteraction,
		a.dataInstruction,
	)
	if err != nil {
		return err
	}
	return nil
}

func NewAddAvailableInteraction(
	availableInteraction interaction.DataAvailableInteraction,
	dataInstruction instruction.DataInstruction,
) (*AddAvailableInteraction, error) {
	return &AddAvailableInteraction{
		availableInteraction: availableInteraction,
		dataInstruction:      dataInstruction,
	}, nil
}
