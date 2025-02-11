package phase

import (
	"github.com/big-smiles/golang-boardgames/pkg/interaction"
)

type IInitializationContext interface {
	GetManagerInteraction() *interaction.ManagerInteraction
}
