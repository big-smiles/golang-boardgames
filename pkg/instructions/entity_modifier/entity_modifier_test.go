package instructionEntityModifier

import (
	"errors"
	"github.com/big-smiles/boardgame-golang/pkg/entity"
	"github.com/big-smiles/boardgame-golang/pkg/game"
	"github.com/big-smiles/boardgame-golang/pkg/instruction"
	instructionControl "github.com/big-smiles/boardgame-golang/pkg/instructions/control"
	instructionEntity "github.com/big-smiles/boardgame-golang/pkg/instructions/entity"
	instructionOutput "github.com/big-smiles/boardgame-golang/pkg/instructions/output"
	"github.com/big-smiles/boardgame-golang/pkg/interaction"
	"github.com/big-smiles/boardgame-golang/pkg/output"
	"github.com/big-smiles/boardgame-golang/pkg/phase"
	"github.com/big-smiles/boardgame-golang/pkg/player"
	"github.com/big-smiles/boardgame-golang/pkg/resolve_value/constant"
	ValueModifierCommon "github.com/big-smiles/boardgame-golang/pkg/value_modifier/common"
	"testing"
)

// TODO: add instruction to this test
func TestAddEntityModifier(t *testing.T) {
	var playerId player.Id = "player_1"
	var nameEntity entity.NameEntityId = "entity_1"
	var nameProperty entity.NamePropertyId[bool] = "property_1"
	var nameDataEntity entity.NameDataEntity = "data_entity_1"
	var namePhase phase.NamePhase = "phase_1"
	calledTimes := 0
	callback := func(output *output.Game) {
		calledTimes++
		if output == nil {
			t.Fatal("ManagerGame is nil")
		}
		if len(output.Entities) != 1 {
			t.Fatal("expected 1 entities")
		}
		var ent entity.OutputEntity
		for _, e := range output.Entities {
			if e.Id == 2 {
				ent = e
			}
		}
		switch calledTimes {
		case 1:
			if ent.Properties.BoolProperties[1] != false {
				t.Fatal("before the modifier the value should be false")
			}
		case 2:
			if ent.Properties.BoolProperties[1] != true {
				t.Fatal("after the modifier the value should be true")
			}
		default:
			t.Fatal("calledTimes should be 1 or 2")
		}
	}
	var callbackInteraction interaction.Callback = func([]interaction.OutputInteraction) {
		t.Log("callback interaction called")
	}
	stringValueResolver, err := resolveValueConstant.NewResolveConstant[entity.NameEntityId](nameEntity)
	if err != nil {
		t.Fatal(err)
	}

	id, err := entity.NewDataId(stringValueResolver)
	if err != nil {
		t.Fatal(err)
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
		t.Fatal(err)
	}
	entityData, err := entity.NewDataEntity(*id, *dataProperties)
	if err != nil {
		t.Fatal(err)
	}
	libraryDataEntities := make(entity.LibraryDataEntity, 1)
	libraryDataEntities[nameDataEntity] = *entityData

	dataCreateEntity, err := instructionEntity.NewDataInstructionCreateEntity(nameDataEntity)
	if err != nil {
		t.Fatal(err)
	}
	dataSendOutput1, err := instructionOutput.NewDataInstructionSendOutput()
	if err != nil {
		t.Fatal(err)
	}
	dataSendOutput2, err := instructionOutput.NewDataInstructionSendOutput()
	if err != nil {
		t.Fatal(err)
	}

	boolModifiers := make(entity.MapDataModifierProperties[bool], 1)
	valueResolver, err := resolveValueConstant.NewResolveConstant[bool](true)
	if err != nil {
		t.Fatal(err)
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
		t.Fatal(err)
	}

	dataEntityModifier, err := entity.NewDataEntityModifier(*dataPropertiesModifier)
	if err != nil {
		t.Fatal(err)
	}

	dataInstructionAddModifier, err := NewDataInstructionAddEntityModifier(*id, *dataEntityModifier)
	if err != nil {
		t.Fatal(err)
	}
	arr := make([]instruction.DataInstruction, 4)
	arr[0] = *dataCreateEntity
	arr[1] = *dataSendOutput1
	arr[2] = *dataInstructionAddModifier
	arr[3] = *dataSendOutput2

	dArray, err := instructionControl.NewDataInstructionArray(arr)
	if err != nil {
		t.Fatal(err)
	}
	stage1, err := phase.NewDataStage(dArray)
	if err != nil {
		t.Fatal(err)
	}

	turn1, err := phase.NewDataTurn([]player.Id{playerId}, []phase.DataStage{*stage1})
	if err != nil {
		t.Fatal(err)
	}

	p, err := phase.NewDataPhase([]phase.DataTurn{*turn1})
	if err != nil {
		t.Fatal(err)
	}
	libraryPhase := make(phase.LibraryPhase, 1)
	libraryPhase[namePhase] = *p
	players := []player.Id{"player1"}
	gameData, err := game.NewDataGame(
		libraryDataEntities,
		libraryPhase,
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
		var errorNoNextPhase phase.ErrorNoNextPhase
		if errors.As(err, &errorNoNextPhase) {
			t.Log(err.Error())
		} else {
			t.Fatal(err)
		}
	}
	if calledTimes != 2 {
		t.Fatal("expected 1 calledTimes")
	}
}
