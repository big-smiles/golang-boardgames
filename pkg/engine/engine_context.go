package engine

import (
	"github.com/big-smiles/golang-boardgames/pkg/entity"
	"github.com/big-smiles/golang-boardgames/pkg/instruction"
	"github.com/big-smiles/golang-boardgames/pkg/interaction"
	"github.com/big-smiles/golang-boardgames/pkg/output"
	"github.com/big-smiles/golang-boardgames/pkg/phase"
	"github.com/big-smiles/golang-boardgames/pkg/player"
)

type EngineContext struct {
	performer                 *instruction.Performer
	managerInstruction        *instruction.ManagerInstruction
	managerEntity             *entity.ManagerEntity
	managerEntityId           *entity.ManagerEntityId
	managerEntityData         *entity.ManagerData
	managerPropertyId         *entity.ManagerPropertyId
	managerOutput             *output.ManagerOutput
	ManagerPhase              *phase.ManagerPhase
	ManagerInteraction        *interaction.ManagerInteraction
	managerTriggerInstruction *instruction.ManagerTriggerInstruction
	managerPlayer             *player.ManagerPlayer
}

// NewEngineContext Returns a new EngineContext
//
// # Params
//
// entityData entity.LibraryDataEntity the set of entityData that will be used to create
// entities during this engine
//
// callback output.Callback the callback to be called every time a new output is issued
//
// phases phase.LibraryPhase the set of phases that will be available during the engine
// firstPhase phase.NamePhase the phase that will be loaded when the engine begins
func NewEngineContext(
	entityData entity.LibraryDataEntity,
	callbackOutput output.Callback,
	phases phase.LibraryPhase,
	firstPhase phase.NamePhase,
	callbackInteraction interaction.Callback,
	players []player.Id,
) (*EngineContext, error) {
	p, err := instruction.NewPerformer()
	if err != nil {
		return nil, err
	}

	mi, err := instruction.NewManagerInstruction()
	if err != nil {
		return nil, err
	}
	managerTriggerInstruction, err := instruction.NewManagerTriggerInstruction()
	if err != nil {
		return nil, err
	}

	managerEntity, err := entity.NewManagerEntity()
	if err != nil {
		return nil, err
	}

	managerEntityData, err := entity.NewManagerData(entityData)
	if err != nil {
		return nil, err
	}

	managerEntityId, err := entity.NewManagerEntityId()
	if err != nil {
		return nil, err
	}

	managerPropertyId, err := entity.NewManagerPropertyId()
	if err != nil {
		return nil, err
	}

	managerOutput, err := output.NewManagerOutput(callbackOutput)
	if err != nil {
		return nil, err
	}

	managerPhase, err := phase.NewManagerPhase(phases, firstPhase)
	if err != nil {
		return nil, err
	}

	managerInteraction, err := interaction.NewManagerInteraction(callbackInteraction)
	if err != nil {
		return nil, err
	}
	managerPlayer, err := player.NewManagerPlayer(players)
	if err != nil {
		return nil, err
	}

	engineContext := EngineContext{
		performer:                 p,
		managerInstruction:        mi,
		managerEntity:             managerEntity,
		managerEntityData:         managerEntityData,
		managerEntityId:           managerEntityId,
		managerPropertyId:         managerPropertyId,
		managerOutput:             managerOutput,
		ManagerPhase:              managerPhase,
		ManagerInteraction:        managerInteraction,
		managerTriggerInstruction: managerTriggerInstruction,
		managerPlayer:             managerPlayer,
	}
	err = engineContext.initialize()
	if err != nil {
		return nil, err
	}
	return &engineContext, nil
}

func (e EngineContext) initialize() error {
	err := e.managerInstruction.Initialize(e)
	if err != nil {
		return err
	}

	err = e.managerTriggerInstruction.Initialize(e)
	if err != nil {
		return err
	}

	err = e.performer.Initialize(e)
	if err != nil {
		return err
	}

	err = e.managerEntity.Initialize(e)
	if err != nil {
		return err
	}

	err = e.managerEntityId.Initialize(e)
	if err != nil {
		return err
	}

	err = e.managerEntityData.Initialize(e)
	if err != nil {
		return err
	}

	err = e.managerPropertyId.Initialize(e)
	if err != nil {
		return err
	}

	err = e.managerOutput.Initialize(e)
	if err != nil {
		return err
	}

	err = e.ManagerPhase.Initialize(e)
	if err != nil {
		return err
	}

	err = e.ManagerInteraction.Initialize(e)
	if err != nil {
		return err
	}

	err = e.managerPlayer.Initialize(e)
	if err != nil {
		return err
	}

	return nil
}

// GetPerformer we declare getters to fix cyclic dependencies by passing the ctx as
// an interface declared on the manager packages
func (e EngineContext) GetPerformer() *instruction.Performer {
	return e.performer
}

// GetManagerEntity we declare getters to fix cyclic dependencies by passing the ctx as
// an interface declared on the manager packages
func (e EngineContext) GetManagerEntity() *entity.ManagerEntity {
	return e.managerEntity
}

// GetManagerOutput we declare getters to fix cyclic dependencies by passing the ctx as
// an interface declared on the manager packages
func (e EngineContext) GetManagerOutput() *output.ManagerOutput {
	return e.managerOutput
}

// GetManagerEntityId we declare getters to fix cyclic dependencies by passing the ctx as
// an interface declared on the manager packages
func (e EngineContext) GetManagerEntityId() *entity.ManagerEntityId {
	return e.managerEntityId
}

// GetManagerPropertyId we declare getters to fix cyclic dependencies by passing the ctx as
// an interface declared on the manager packages
func (e EngineContext) GetManagerPropertyId() *entity.ManagerPropertyId {
	return e.managerPropertyId
}

// GetManagerEntityData we declare getters to fix cyclic dependencies by passing the ctx as
// an interface declared on the manager packages
func (e EngineContext) GetManagerEntityData() *entity.ManagerData {
	return e.managerEntityData
}

// GetManagerInstruction we declare getters to fix cyclic dependencies by passing the ctx as
// an interface declared on the manager packages
func (e EngineContext) GetManagerInstruction() *instruction.ManagerInstruction {
	return e.managerInstruction
}

// GetITriggerInstruction we declare getters to fix cyclic dependencies by passing the ctx as
// an interface declared on the manager packages
// ITriggerInstruction is an interface on the interaction package
func (e EngineContext) GetITriggerInstruction() interaction.ITriggerInstruction {
	return e.managerTriggerInstruction
}

// GetManagerInteraction we declare getters to fix cyclic dependencies by passing the ctx as
// an interface declared on the manager packages
func (e EngineContext) GetManagerInteraction() *interaction.ManagerInteraction {
	return e.ManagerInteraction
}

// GetManagerTriggerInstruction we declare getters to fix cyclic dependencies by passing the ctx as
// an interface declared on the manager packages
func (e EngineContext) GetManagerTriggerInstruction() *instruction.ManagerTriggerInstruction {
	return e.managerTriggerInstruction
}

// GetManagerPlayer we declare getters to fix cyclic dependencies by passing the ctx as
// an interface declared on the manager packages
func (e EngineContext) GetManagerPlayer() *player.ManagerPlayer {
	return e.managerPlayer
}
