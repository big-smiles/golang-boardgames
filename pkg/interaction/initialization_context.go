package interaction

import (
	"github.com/big-smiles/golang-boardgames/pkg/entity"
	"github.com/big-smiles/golang-boardgames/pkg/player"
)

type IInitializationContext interface {
	GetITriggerInstruction() ITriggerInstruction
	GetManagerPropertyId() *entity.ManagerPropertyId
	GetManagerPlayer() *player.ManagerPlayer
}
