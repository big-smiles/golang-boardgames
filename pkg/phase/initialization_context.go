package phase

import (
	"github.com/big-smiles/boardgame-golang/pkg/instruction"
	"github.com/big-smiles/boardgame-golang/pkg/interaction"
)

type IInitializationContext interface {
	GetManagerInstruction() *instruction.ManagerInstruction
	GetManagerInteraction() *interaction.ManagerInteraction
}
