package instruction

import (
	"github.com/big-smiles/boardgame-golang/pkg/entity"
	"github.com/big-smiles/boardgame-golang/pkg/interaction"
	"github.com/big-smiles/boardgame-golang/pkg/output"
)

type InitializationContext interface {
	GetPerformer() *Performer
	GetManagerEntity() *entity.ManagerEntity
	GetManagerInstruction() *ManagerInstruction
	GetManagerEntityId() *entity.ManagerEntityId
	GetManagerPropertyId() *entity.ManagerPropertyId
	GetManagerOutput() *output.ManagerOutput
	GetManagerEntityData() *entity.ManagerData
	GetManagerInteraction() *interaction.ManagerInteraction
	GetManagerTriggerInstruction() *ManagerTriggerInstruction
}
