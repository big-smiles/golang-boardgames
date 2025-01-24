package interaction

import "github.com/big-smiles/boardgame-golang/pkg/instruction"

type DataClearAvailableInteraction struct {
}

func (d DataClearAvailableInteraction) NewFromThisData() (instruction.Instruction, error) {
	return NewClearAvailableInteraction()
}

func NewDataClearAvailableInteraction() (*DataClearAvailableInteraction, error) {
	return &DataClearAvailableInteraction{}, nil
}
