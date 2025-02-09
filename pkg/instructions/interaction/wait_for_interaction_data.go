package interaction

import "github.com/big-smiles/golang-boardgames/pkg/instruction"

type DataWaitForInteractionData struct {
}

func (d DataWaitForInteractionData) NewFromThisData() (instruction.Instruction, error) {
	return NewWaitForInteraction()
}

func NewDataWaitForInteractionData() *DataWaitForInteractionData {
	return &DataWaitForInteractionData{}
}
