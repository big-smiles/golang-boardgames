package interaction

import (
	"github.com/big-smiles/boardgame-golang/pkg/entity"
	"github.com/big-smiles/boardgame-golang/pkg/player"
)

type IInitializationContext interface {
	GetITriggerInstruction() ITriggerInstruction
	GetManagerPropertyId() *entity.ManagerPropertyId
	GetManagerPlayer() *player.ManagerPlayer
}
