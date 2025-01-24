package instruction

import (
	"github.com/big-smiles/boardgame-golang/pkg/entity"
	"github.com/big-smiles/boardgame-golang/pkg/interaction"
)

type Interaction struct {
	managerInteraction        *interaction.ManagerInteraction
	managerTriggerInstruction *ManagerTriggerInstruction
}

func NewPerformerInteraction() (*Interaction, error) {
	return &Interaction{}, nil
}
func (p *Interaction) Initialize(ctx InitializationContext) error {
	p.managerInteraction = ctx.GetManagerInteraction()
	p.managerTriggerInstruction = ctx.GetManagerTriggerInstruction()
	return nil
}

func (p *Interaction) WaitForInteraction() error {
	err := p.managerInteraction.WaitForInteraction()
	if err != nil {
		return err
	}
	return nil
}
func (p *Interaction) AddAvailableInteraction(
	executionVariable entity.Entity,
	availableInteraction interaction.DataAvailableInteraction,
	dataInstruction DataInstruction,
) error {
	id, err := p.managerTriggerInstruction.AddInstructionToTrigger(dataInstruction)
	if err != nil {
		return err
	}
	err = p.managerInteraction.AddAvailableInteraction(executionVariable, availableInteraction, id)
	if err != nil {
		return err
	}
	return nil
}
func (p *Interaction) ClearAvailableInteraction() error {
	err := p.managerInteraction.ClearAvailableInteraction()
	if err != nil {
		return err
	}
	return nil
}
