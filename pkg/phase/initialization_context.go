package phase

import (
	"github.com/big-smiles/golang-boardgames/pkg/instruction"
	"github.com/big-smiles/golang-boardgames/pkg/interaction"
)

type IInitializationContext interface {
	GetManagerInstruction() *instruction.ManagerInstruction
	GetManagerInteraction() *interaction.ManagerInteraction
}
