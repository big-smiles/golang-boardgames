package interaction

import "github.com/big-smiles/boardgame-golang/pkg/instruction"

type DataWaitForInteractionData struct {
}

func (d DataWaitForInteractionData) NewFromThisData() (instruction.Instruction, error) {
	return NewWaitForInteraction()
}

func NewDataWaitForInteractionData() (*DataWaitForInteractionData, error) {
	return &DataWaitForInteractionData{}, nil
}
