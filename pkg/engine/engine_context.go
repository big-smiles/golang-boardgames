package engine

import (
	"github.com/big-smiles/golang-boardgames/pkg/entity"
	"github.com/big-smiles/golang-boardgames/pkg/instruction"
	"github.com/big-smiles/golang-boardgames/pkg/interaction"
	"github.com/big-smiles/golang-boardgames/pkg/output"
	"github.com/big-smiles/golang-boardgames/pkg/phase"
	"github.com/big-smiles/golang-boardgames/pkg/phaseData"
	"github.com/big-smiles/golang-boardgames/pkg/player"
)

type Context struct {
	performer                 *instruction.Performer
	managerInstruction        *instruction.ManagerInstruction
	managerEntity             *entity.ManagerEntity
	managerEntityId           *entity.ManagerEntityId
	managerEntityData         *entity.ManagerData
	managerPropertyId         *entity.ManagerPropertyId
	managerOutput             *output.ManagerOutput
	managerPhase              *phase.ManagerPhase
	ManagerInteraction        *interaction.ManagerInteraction
	managerTriggerInstruction *instruction.ManagerTriggerInstruction
	managerPlayer             *player.ManagerPlayer
}

// NewContext Returns a new Context
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
func NewContext(
	entityData entity.LibraryDataEntity,
	callbackOutput output.Callback,
	phases []phaseData.DataPhase,
	firstPhase phase.NamePhase,
	callbackInteraction interaction.Callback,
	players []player.Id,
) (*Context, error) {
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

	managerPhase, err := phase.NewManagerPhase()
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

	engineContext := Context{
		performer:                 p,
		managerInstruction:        mi,
		managerEntity:             managerEntity,
		managerEntityData:         managerEntityData,
		managerEntityId:           managerEntityId,
		managerPropertyId:         managerPropertyId,
		managerOutput:             managerOutput,
		managerPhase:              managerPhase,
		ManagerInteraction:        managerInteraction,
		managerTriggerInstruction: managerTriggerInstruction,
		managerPlayer:             managerPlayer,
	}
	err = engineContext.initialize()
	if err != nil {
		return nil, err
	}
	err = engineContext.loadPhase(phases, firstPhase)
	if err != nil {
		return nil, err
	}
	return &engineContext, nil
}

func (e *Context) initialize() error {
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

	err = e.managerPhase.Initialize(e)
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
func (e *Context) loadPhase(
	phases []phaseData.DataPhase,
	firstPhase phase.NamePhase,
) error {
	libraryPhase := make(phase.LibraryPhase)
	for _, pd := range phases {
		turns := make([]phase.Turn, len(pd.Turns))
		for i, turn := range pd.Turns {
			stages := make([]phase.Stage, len(turn.Stages))
			for j, stage := range turn.Stages {
				callback := e.getStageCallback(stage.Instructions)
				stages[j] = *phase.NewStage(callback)
			}
			turns[i] = *phase.NewTurn(turn.Name, turn.ActivePlayers, stages)
		}
		p := phase.NewPhase(pd.Name, turns)
		libraryPhase[pd.Name] = *p
	}
	err := e.managerPhase.LoadPhases(libraryPhase, firstPhase)
	if err != nil {
		return err
	}
	return nil
}
func (e *Context) getStageCallback(
	dataInstruction instruction.DataInstruction,
) phase.CallbackInstructionExecute {
	return func() error {
		i, err := dataInstruction.NewFromThisData()
		if err != nil {
			return err
		}
		err = e.managerInstruction.AddInstruction(i, nil)
		if err != nil {
			return err
		}
		return nil
	}
}
func (e *Context) Next() error {
	return e.managerPhase.Next()
}

func (e *Context) GetActivePlayerProvider() output.ActivePlayerProvider {
	//TODO implement me
	return e.managerPhase
}

// GetPerformer we declare getters to fix cyclic dependencies by passing the ctx as
// an interface declared on the manager packages
func (e *Context) GetPerformer() *instruction.Performer {
	return e.performer
}

// GetManagerEntity we declare getters to fix cyclic dependencies by passing the ctx as
// an interface declared on the manager packages
func (e *Context) GetManagerEntity() *entity.ManagerEntity {
	return e.managerEntity
}

// GetManagerOutput we declare getters to fix cyclic dependencies by passing the ctx as
// an interface declared on the manager packages
func (e *Context) GetManagerOutput() *output.ManagerOutput {
	return e.managerOutput
}

// GetManagerEntityId we declare getters to fix cyclic dependencies by passing the ctx as
// an interface declared on the manager packages
func (e *Context) GetManagerEntityId() *entity.ManagerEntityId {
	return e.managerEntityId
}

// GetManagerPropertyId we declare getters to fix cyclic dependencies by passing the ctx as
// an interface declared on the manager packages
func (e *Context) GetManagerPropertyId() *entity.ManagerPropertyId {
	return e.managerPropertyId
}

// GetManagerEntityData we declare getters to fix cyclic dependencies by passing the ctx as
// an interface declared on the manager packages
func (e *Context) GetManagerEntityData() *entity.ManagerData {
	return e.managerEntityData
}

// GetManagerInstruction we declare getters to fix cyclic dependencies by passing the ctx as
// an interface declared on the manager packages
func (e *Context) GetManagerInstruction() *instruction.ManagerInstruction {
	return e.managerInstruction
}

// GetITriggerInstruction we declare getters to fix cyclic dependencies by passing the ctx as
// an interface declared on the manager packages
// ITriggerInstruction is an interface on the interaction package
func (e *Context) GetITriggerInstruction() interaction.ITriggerInstruction {
	return e.managerTriggerInstruction
}

// GetManagerInteraction we declare getters to fix cyclic dependencies by passing the ctx as
// an interface declared on the manager packages
func (e *Context) GetManagerInteraction() *interaction.ManagerInteraction {
	return e.ManagerInteraction
}

// GetManagerTriggerInstruction we declare getters to fix cyclic dependencies by passing the ctx as
// an interface declared on the manager packages
func (e *Context) GetManagerTriggerInstruction() *instruction.ManagerTriggerInstruction {
	return e.managerTriggerInstruction
}

// GetManagerPhase we declare getters to fix cyclic dependencies by passing the ctx as
// an interface declared on the manager packages
func (e *Context) GetManagerPhase() *phase.ManagerPhase {
	return e.managerPhase
}

// GetManagerPlayer we declare getters to fix cyclic dependencies by passing the ctx as
// an interface declared on the manager packages
func (e *Context) GetManagerPlayer() *player.ManagerPlayer {
	return e.managerPlayer
}
