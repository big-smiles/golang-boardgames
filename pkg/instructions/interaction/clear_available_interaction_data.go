package interaction

import "github.com/big-smiles/golang-boardgames/pkg/instruction"

type DataClearAvailableInteraction struct {
}

func (d DataClearAvailableInteraction) NewFromThisData() (instruction.Instruction, error) {
	return NewClearAvailableInteraction()
}

func NewDataClearAvailableInteraction() *DataClearAvailableInteraction {
	return &DataClearAvailableInteraction{}
}
