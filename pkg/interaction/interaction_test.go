package interaction_test

import (
	"errors"
	"github.com/big-smiles/boardgame-golang/pkg/entity"
	"github.com/big-smiles/boardgame-golang/pkg/game"
	"github.com/big-smiles/boardgame-golang/pkg/instruction"
	instructionControl "github.com/big-smiles/boardgame-golang/pkg/instructions/control"
	instructionEntity "github.com/big-smiles/boardgame-golang/pkg/instructions/entity"
	instructionEntityModifier "github.com/big-smiles/boardgame-golang/pkg/instructions/entity_modifier"
	instructionInteraction "github.com/big-smiles/boardgame-golang/pkg/instructions/interaction"
	instructionOutput "github.com/big-smiles/boardgame-golang/pkg/instructions/output"
	"github.com/big-smiles/boardgame-golang/pkg/interaction"
	"github.com/big-smiles/boardgame-golang/pkg/output"
	"github.com/big-smiles/boardgame-golang/pkg/phase"
	"github.com/big-smiles/boardgame-golang/pkg/player"
	resolveValueConstant "github.com/big-smiles/boardgame-golang/pkg/resolve_value/constant"
	ValueModifierCommon "github.com/big-smiles/boardgame-golang/pkg/value_modifier/common"
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
	var interactionToSelect []interaction.SelectedInteraction = make([]interaction.SelectedInteraction, 1)
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
	stringValueResolver, err := resolveValueConstant.NewResolveConstant[entity.NameEntityId](nameEntity)
	if err != nil {
		return nil, nil, err
	}

	id, err := entity.NewDataId(stringValueResolver)
	if err != nil {
		return nil, nil, err
	}
	stringValueResolver2, err := resolveValueConstant.NewResolveConstant[entity.NameEntityId](nameEntity2)
	if err != nil {
		return nil, nil, err
	}

	id2, err := entity.NewDataId(stringValueResolver2)
	if err != nil {
		return nil, nil, err
	}

	namesBoolProperties := make([]entity.NamePropertyId[bool], 1)
	namesStringProperties := make([]entity.NamePropertyId[string], 0)
	namesIntProperties := make([]entity.NamePropertyId[int], 0)
	namesEntityIdProperties := make([]entity.NamePropertyId[entity.Id], 0)
	namesArrayEntityIdProperties := make([]entity.NamePropertyId[[]entity.Id], 0)

	namesBoolProperties[0] = nameProperty

	dataProperties, err := entity.NewDataProperties(
		namesBoolProperties,
		namesStringProperties,
		namesEntityIdProperties,
		namesIntProperties,
		namesArrayEntityIdProperties,
	)
	if err != nil {
		return nil, nil, err
	}
	entityData, err := entity.NewDataEntity(*id, *dataProperties)
	entityData2, err := entity.NewDataEntity(*id2, *dataProperties)
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

) (*phase.LibraryPhase, error) {

	dataCreateEntity, err := instructionEntity.NewDataInstructionCreateEntity(nameDataEntity)
	if err != nil {
		return nil, err
	}
	dataCreateEntity2, err := instructionEntity.NewDataInstructionCreateEntity(nameDataEntity2)
	if err != nil {
		return nil, err
	}
	dataSendOutput1, err := instructionOutput.NewDataInstructionSendOutput()
	if err != nil {
		return nil, err
	}
	dataSendOutput2, err := instructionOutput.NewDataInstructionSendOutput()
	if err != nil {
		return nil, err
	}

	boolModifiers := make(entity.MapDataModifierProperties[bool], 1)
	valueResolver, err := resolveValueConstant.NewResolveConstant[bool](true)
	if err != nil {
		return nil, err
	}
	boolModifiers[nameProperty], err = ValueModifierCommon.NewDataModifierSetValue(valueResolver)
	dataPropertiesModifier, err := entity.NewDataPropertiesModifier(
		nil,
		nil,
		&boolModifiers,
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}

	dataEntityModifier, err := entity.NewDataEntityModifier(*dataPropertiesModifier)
	if err != nil {
		return nil, err
	}
	resolveFromSelectedId, err := resolveValueConstant.NewResolveValueFromVariable[[]entity.Id](instruction.SELECTED_ENTITIES)
	dataInstructionAddModifier, err := instructionEntityModifier.NewDataInstructionAddEntityModifierWithResolvedTarget(
		resolveFromSelectedId,
		*dataEntityModifier,
	)
	if err != nil {
		return nil, err
	}

	availableEntityIds, err := resolveValueConstant.NewResolveValueFromVariable(variablePropertyName)
	if err != nil {
		return nil, err
	}
	dataAvailableInteraction, err := interaction.NewDataAvailableInteraction(
		playerId,
		availableEntityIds,
		1,
		2,
	)
	if err != nil {
		return nil, err
	}

	var predicate entity.Predicate = func(
		executionVariable entity.Entity,
		e entity.Entity,
	) (bool, error) {
		return e.Name == nameEntity || e.Name == nameEntity2, nil
	}
	dataFilterEntity, err := instructionEntity.NewDataInstructionFilterEntities(
		predicate,
		variablePropertyName,
	)
	if err != nil {
		return nil, err
	}

	arr2 := make([]instruction.DataInstruction, 2)
	arr2[0] = *dataInstructionAddModifier
	arr2[1] = *dataSendOutput2
	instructionArray2, err := instructionControl.NewDataInstructionArray(arr2)
	if err != nil {
		return nil, err
	}

	dataInstructionAddAvailableInteraction, err := instructionInteraction.NewDataAvailableInteractionData(
		*dataAvailableInteraction,
		instructionArray2,
	)
	if err != nil {
		return nil, err
	}

	dataInstructionWaitForInteraction, err := instructionInteraction.NewDataWaitForInteractionData()
	if err != nil {
		return nil, err
	}

	arr := make([]instruction.DataInstruction, 6)
	arr[0] = *dataCreateEntity
	arr[1] = *dataCreateEntity2
	arr[2] = *dataSendOutput1
	arr[3] = *dataFilterEntity
	arr[4] = *dataInstructionAddAvailableInteraction
	arr[5] = *dataInstructionWaitForInteraction

	instructionArray1, err := instructionControl.NewDataInstructionArray(arr)
	if err != nil {
		return nil, err
	}

	stage1, err := phase.NewDataStage(instructionArray1)
	if err != nil {
		return nil, err
	}

	turn1, err := phase.NewDataTurn([]player.Id{playerId}, []phase.DataStage{*stage1})
	if err != nil {
		return nil, err
	}

	p1, err := phase.NewDataPhase([]phase.DataTurn{*turn1})
	if err != nil {
		return nil, err
	}
	libraryPhase := make(phase.LibraryPhase, 1)
	libraryPhase[namePhase1] = *p1
	return &libraryPhase, nil
}
