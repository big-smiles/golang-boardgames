package phase

import (
	"github.com/big-smiles/golang-boardgames/pkg/instruction"
)

type DataStage struct {
	instructions instruction.DataInstruction
}

func NewDataStage(instructions instruction.DataInstruction) (*DataStage, error) {
	return &DataStage{
		instructions: instructions,
	}, nil
}
