package interaction_test

import (
	"errors"
	"github.com/big-smiles/golang-boardgames/pkg/entity"
	"github.com/big-smiles/golang-boardgames/pkg/game"
	"github.com/big-smiles/golang-boardgames/pkg/instruction"
	instructionControl "github.com/big-smiles/golang-boardgames/pkg/instructions/control"
	instructionEntity "github.com/big-smiles/golang-boardgames/pkg/instructions/entity"
	instructionEntityModifier "github.com/big-smiles/golang-boardgames/pkg/instructions/entity_modifier"
	instructionInteraction "github.com/big-smiles/golang-boardgames/pkg/instructions/interaction"
	instructionOutput "github.com/big-smiles/golang-boardgames/pkg/instructions/output"
	"github.com/big-smiles/golang-boardgames/pkg/interaction"
	"github.com/big-smiles/golang-boardgames/pkg/output"
	"github.com/big-smiles/golang-boardgames/pkg/phase"
	"github.com/big-smiles/golang-boardgames/pkg/phaseData"
	"github.com/big-smiles/golang-boardgames/pkg/player"
	resolveValueConstant "github.com/big-smiles/golang-boardgames/pkg/resolve_value/constant"
	ValueModifierCommon "github.com/big-smiles/golang-boardgames/pkg/value_modifier/common"
	"testing"
)

func TestInteractions(t *testing.T) {
	var playerId player.Id = "player_1"
	var nameEntity entity.NameEntityId = "entity_1"
	var nameEntity2 entity.NameEntityId = "entity_2"
	var nameProperty entity.NamePropertyId[bool] = "property_1"
	var variablePropertyName entity.NamePropertyId[[]entity.Id] = "variable_property_1"
	var nameDataEntity entity.NameDataEntity = "data_entity_1"
	var nameDataEntity2 entity.NameDataEntity = "data_entity_2"
	var namePhase phase.NamePhase = "phase_1"
	calledTimesOutput := 0
	callback := func(output *output.Game) {
		calledTimesOutput++
		if output == nil {
			t.Fatal("ManagerGame is nil")
		}
		if len(output.Entities) != 2 {
			t.Fatal("expected 2 entities ")
		}
		var ent entity.OutputEntity
		for _, e := range output.Entities {
			if e.Name == nameEntity2 {
				ent = e
			}
		}
		switch calledTimesOutput {
		case 1:
			if ent.Properties.BoolProperties[1] != false {
				t.Fatal("before the modifier the value should be false")
			}
		case 2:
			if ent.Properties.BoolProperties[1] != true {
				t.Fatal("after the modifier the value should be true")
			}
		default:
			t.Fatal("calledTimesOutput should be 1 or 2")
		}
	}
	calledTimesInteraction := 0
	var interactionToSelect = make([]interaction.SelectedInteraction, 1)
	var callbackInteraction interaction.Callback = func(outputInteraction []interaction.OutputInteraction) {
		t.Log("callback interaction called")
		calledTimesInteraction++
		id := outputInteraction[0].Id
		var err error
		selectedEntities := make([]entity.Id, 1)
		selectedEntities[0] = outputInteraction[0].AvailableEntities[1]
		a, err := interaction.NewSelectedInteraction(id, playerId, selectedEntities)
		if err != nil {
			t.Fatal(err)
		}
		interactionToSelect[0] = *a
	}

	libraryDataEntities, _, err := createDataEntityLibrary(
		nameEntity,
		nameEntity2,
		nameProperty,
		nameDataEntity,
		nameDataEntity2,
	)
	if err != nil {
		t.Fatal(err)
	}
	libraryPhase, err := createPhaseLibrary(
		nameEntity,
		nameEntity2,
		nameDataEntity,
		nameDataEntity2,
		nameProperty,
		namePhase,
		variablePropertyName,
		playerId,
	)
	if err != nil {
		t.Fatal(err)
	}
	players := []player.Id{playerId}
	gameData, err := game.NewDataGame(
		*libraryDataEntities,
		*libraryPhase,
		namePhase,
		players,
	)
	if err != nil {
		t.Fatal(err)
	}

	g, err := game.NewGame(*gameData, callback, callbackInteraction)
	if err != nil {
		t.Fatal(err)
	}
	err = g.Start()
	if err != nil {
		t.Fatal(err)
	}
	if calledTimesOutput != 1 {
		t.Fatal("expected 1 calledTimesOutput")
	}
	if calledTimesInteraction != 1 {
		t.Fatal("expected 1 calledTimesInteraction")
	}
	err = g.SelectInteraction(interactionToSelect)
	if err != nil {
		var errorNoNextPhase phase.ErrorNoNextPhase
		if errors.As(err, &errorNoNextPhase) {
			t.Log(err.Error())
		} else {
			t.Fatal(err)
		}
	}
	if calledTimesOutput != 2 {
		t.Fatal("expected 1 calledTimesOutput")
	}
	if calledTimesInteraction != 1 {
		t.Fatal("expected 1 calledTimesInteraction")
	}
}

func createDataEntityLibrary(
	nameEntity entity.NameEntityId,
	nameEntity2 entity.NameEntityId,
	nameProperty entity.NamePropertyId[bool],
	nameDataEntity entity.NameDataEntity,
	nameDataEntity2 entity.NameDataEntity,
) (*entity.LibraryDataEntity, *entity.DataId, error) {
	stringValueResolver := resolveValueConstant.NewResolveConstant[entity.NameEntityId](nameEntity)

	id, err := entity.NewDataId(stringValueResolver)
	if err != nil {
		return nil, nil, err
	}
	stringValueResolver2 := resolveValueConstant.NewResolveConstant[entity.NameEntityId](nameEntity2)

	id2, err := entity.NewDataId(stringValueResolver2)
	if err != nil {
		return nil, nil, err
	}

	dataProperties := entity.DataProperties{
		BoolProperties: []entity.NamePropertyId[bool]{
			nameProperty,
		}}
	entityData, err := entity.NewDataEntity(*id, dataProperties)
	entityData2, err := entity.NewDataEntity(*id2, dataProperties)
	if err != nil {
		return nil, nil, err
	}
	libraryDataEntities := make(entity.LibraryDataEntity, 2)
	libraryDataEntities[nameDataEntity] = *entityData
	libraryDataEntities[nameDataEntity2] = *entityData2
	return &libraryDataEntities, id, nil
}

func createPhaseLibrary(
	nameEntity entity.NameEntityId,
	nameEntity2 entity.NameEntityId,
	nameDataEntity entity.NameDataEntity,
	nameDataEntity2 entity.NameDataEntity,
	nameProperty entity.NamePropertyId[bool],
	namePhase1 phase.NamePhase,
	variablePropertyName entity.NamePropertyId[[]entity.Id],
	playerId player.Id,

) (*[]phaseData.DataPhase, error) {

	boolModifiers := make(entity.MapDataModifierProperties[bool], 1)
	valueResolver := resolveValueConstant.NewResolveConstant[bool](true)
	bolModifier, err := ValueModifierCommon.NewDataModifierSetValue(valueResolver)
	boolModifiers[nameProperty] = bolModifier

	dataPropertiesModifier := entity.DataPropertiesModifier{
		BoolModifiers: boolModifiers,
	}

	dataEntityModifier, err := entity.NewDataEntityModifier(dataPropertiesModifier)
	if err != nil {
		return nil, err
	}

	resolveFromSelectedId := resolveValueConstant.NewResolveValueFromVariable[[]entity.Id](instruction.SelectedEntities)

	availableEntityIds := resolveValueConstant.NewResolveValueFromVariable(variablePropertyName)

	dataAvailableInteraction, err := interaction.NewDataAvailableInteraction(
		playerId,
		availableEntityIds,
		1,
		2,
	)
	if err != nil {
		return nil, err
	}

	p1 := []phaseData.DataPhase{{
		Name: namePhase1,
		Turns: []phaseData.DataTurn{
			{
				ActivePlayers: []player.Id{playerId},
				Stages: []phaseData.DataStage{
					{
						Instructions: instructionControl.NewDataInstructionArray(
							instructionEntity.NewDataInstructionCreateEntity(nameDataEntity),
							instructionEntity.NewDataInstructionCreateEntity(nameDataEntity2),
							instructionOutput.NewDataInstructionSendOutput(),
							instructionEntity.NewDataInstructionFilterEntities(
								func(
									executionVariable entity.Entity,
									managerPropertyId *entity.ManagerPropertyId,
									e entity.Entity,
								) (bool, error) {
									return e.Name == nameEntity || e.Name == nameEntity2, nil
								},
								variablePropertyName,
							),
							instructionInteraction.NewDataAvailableInteractionData(
								*dataAvailableInteraction,
								instructionControl.NewDataInstructionArray(
									instructionEntityModifier.NewDataInstructionAddEntityModifierWithResolvedTarget(
										resolveFromSelectedId,
										*dataEntityModifier,
									),
									instructionOutput.NewDataInstructionSendOutput(),
								),
							),
							instructionInteraction.NewDataWaitForInteractionData(),
						),
					},
				},
			},
		},
	}}

	return &p1, nil
}
