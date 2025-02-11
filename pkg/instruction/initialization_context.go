package instruction

import (
	"github.com/big-smiles/golang-boardgames/pkg/entity"
	"github.com/big-smiles/golang-boardgames/pkg/interaction"
	"github.com/big-smiles/golang-boardgames/pkg/output"
	"github.com/big-smiles/golang-boardgames/pkg/phase"
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
	GetManagerPhase() *phase.ManagerPhase
}
