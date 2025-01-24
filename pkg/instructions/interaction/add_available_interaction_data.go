package interaction

import (
	"github.com/big-smiles/golang-boardgames/pkg/instruction"
	"github.com/big-smiles/golang-boardgames/pkg/interaction"
)

type DataAddAvailableInteractionData struct {
	availableInteraction interaction.DataAvailableInteraction
	dataInstruction      instruction.DataInstruction
}

func NewDataAvailableInteractionData(availableInteraction interaction.DataAvailableInteraction,
	dataInstruction instruction.DataInstruction) (*DataAddAvailableInteractionData, error) {
	return &DataAddAvailableInteractionData{
		availableInteraction: availableInteraction,
		dataInstruction:      dataInstruction,
	}, nil
}
func (d DataAddAvailableInteractionData) NewFromThisData() (instruction.Instruction, error) {
	return NewAddAvailableInteraction(d.availableInteraction, d.dataInstruction)
}
