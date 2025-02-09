package instructionEntityModifier

import (
	"errors"
	"github.com/big-smiles/golang-boardgames/pkg/entity"
	"github.com/big-smiles/golang-boardgames/pkg/game"
	instructionControl "github.com/big-smiles/golang-boardgames/pkg/instructions/control"
	instructionEntity "github.com/big-smiles/golang-boardgames/pkg/instructions/entity"
	instructionOutput "github.com/big-smiles/golang-boardgames/pkg/instructions/output"
	"github.com/big-smiles/golang-boardgames/pkg/interaction"
	"github.com/big-smiles/golang-boardgames/pkg/output"
	"github.com/big-smiles/golang-boardgames/pkg/phase"
	"github.com/big-smiles/golang-boardgames/pkg/player"
	"github.com/big-smiles/golang-boardgames/pkg/resolve_value/constant"
	ValueModifierCommon "github.com/big-smiles/golang-boardgames/pkg/value_modifier/common"
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
	stringValueResolver := resolveValueConstant.NewResolveConstant[entity.NameEntityId](nameEntity)

	id, err := entity.NewDataId(stringValueResolver)
	if err != nil {
		t.Fatal(err)
	}

	dataProperties := entity.DataProperties{
		BoolProperties: []entity.NamePropertyId[bool]{
			nameProperty,
		},
	}
	entityData, err := entity.NewDataEntity(*id, dataProperties)
	if err != nil {
		t.Fatal(err)
	}
	libraryDataEntities := make(entity.LibraryDataEntity, 1)
	libraryDataEntities[nameDataEntity] = *entityData

	boolModifiers := make(entity.MapDataModifierProperties[bool], 1)
	valueResolver := resolveValueConstant.NewResolveConstant[bool](true)

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

	p := phase.DataPhase{
		Name: namePhase,
		Turns: []phase.DataTurn{
			{
				Stages: []phase.DataStage{
					{
						Instructions: instructionControl.NewDataInstructionArray(
							instructionEntity.NewDataInstructionCreateEntity(nameDataEntity),
							instructionOutput.NewDataInstructionSendOutput(),
							NewDataInstructionAddEntityModifier(*id, *dataEntityModifier),
							instructionOutput.NewDataInstructionSendOutput(),
						),
					},
				},
				ActivePlayers: []player.Id{playerId},
			},
		},
	}

	libraryPhase := make(phase.LibraryPhase, 1)
	libraryPhase[namePhase] = p
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
