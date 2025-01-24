package instruction

import "github.com/big-smiles/boardgame-golang/pkg/entity"

type wrapperInstruction struct {
	instruction      Instruction
	selectedEntities []entity.Id
}

func newWrapperInstruction(
	instruction Instruction,
	selectedEntities []entity.Id,
) (*wrapperInstruction, error) {
	return &wrapperInstruction{
		instruction:      instruction,
		selectedEntities: selectedEntities,
	}, nil
}
