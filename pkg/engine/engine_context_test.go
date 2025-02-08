package engine_test

import (
	"errors"
	"github.com/big-smiles/golang-boardgames/pkg/entity"
	"github.com/big-smiles/golang-boardgames/pkg/game"
	"github.com/big-smiles/golang-boardgames/pkg/instruction"
	instructionControl "github.com/big-smiles/golang-boardgames/pkg/instructions/control"
	instructionEntity "github.com/big-smiles/golang-boardgames/pkg/instructions/entity"
	instructionOutput "github.com/big-smiles/golang-boardgames/pkg/instructions/output"
	"github.com/big-smiles/golang-boardgames/pkg/interaction"
	"github.com/big-smiles/golang-boardgames/pkg/output"
	"github.com/big-smiles/golang-boardgames/pkg/phase"
	"github.com/big-smiles/golang-boardgames/pkg/player"
	"github.com/big-smiles/golang-boardgames/pkg/resolve_value/constant"
	"testing"
)

// TODO: add instruction to this test
func TestCreateEntity(t *testing.T) {
	calledTimes := 0
	var playerId player.Id = "player_1"
	var nameDataEntity1 entity.NameDataEntity = "data_entity_1"
	var nameEntityId entity.NameEntityId = "entity_1"
	var namePhase phase.NamePhase = "phase_1"
	libraryDataEntity, err := createLibraryDataEntity(nameEntityId, nameDataEntity1)
	if err != nil {
		panic(err)
	}
	callbackOutput := func(output *output.Game) {
		calledTimes++
		if output == nil {
			t.Fatal("ManagerGame is nil")
		}
		if len(output.Entities) != 1 {
			t.Fatal("expected 1 entitys")
		}
		if output.Entities[0].Id != 2 {
			t.Fatal("expected entity to have Id 1")
		}
		if output.Entities[0].Name != nameEntityId {
			t.Fatal("expected entity to have Name 1")
		}
	}
	var callbackInteraction interaction.Callback = func([]interaction.OutputInteraction) {
		t.Log("callback interaction called")
	}
	dCreateEntity, err := instructionEntity.NewDataInstructionCreateEntity(nameDataEntity1)
	if err != nil {
		t.Fatal(err)
	}
	dSendOutput, err := instructionOutput.NewDataInstructionSendOutput()
	if err != nil {
		t.Fatal(err)
	}
	arr := make([]instruction.DataInstruction, 2)
	arr[0] = *dCreateEntity
	arr[1] = *dSendOutput

	dArray, err := instructionControl.NewDataInstructionArray(arr)
	if err != nil {
		t.Fatal(err)
	}

	p := phase.DataPhase{
		Name: namePhase,
		Turns: []phase.DataTurn{
			{
				ActivePlayers: []player.Id{playerId},
				Stages: []phase.DataStage{
					{
						Instructions: dArray,
					},
				},
			},
		},
	}

	phases := make(map[phase.NamePhase]phase.DataPhase, 1)
	phases[namePhase] = p
	players := []player.Id{"player1"}
	dataGame, err := game.NewDataGame(
		*libraryDataEntity,
		phases,
		namePhase,
		players,
	)
	if err != nil {
		t.Fatal(err)
	}
	g, err := game.NewGame(*dataGame, callbackOutput, callbackInteraction)
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

	if calledTimes != 1 {
		t.Fatal("expected 1 calledTimes")
	}
}

func createLibraryDataEntity(nameEntityId entity.NameEntityId, nameDataEntity1 entity.NameDataEntity) (*entity.LibraryDataEntity, error) {
	stringValueResolver, err := resolveValueConstant.NewResolveConstant[entity.NameEntityId](nameEntityId)
	if err != nil {
		return nil, err
	}

	id, err := entity.NewDataId(stringValueResolver)
	if err != nil {
		return nil, err
	}

	namesBoolProperties := make([]entity.NamePropertyId[bool], 0)
	namesStringProperties := make([]entity.NamePropertyId[string], 0)
	namesIntProperties := make([]entity.NamePropertyId[int], 0)
	namesEntityIdProperties := make([]entity.NamePropertyId[entity.Id], 0)
	namesArrayEntityIdProperties := make([]entity.NamePropertyId[[]entity.Id], 0)
	dataProperties, err := entity.NewDataProperties(
		namesBoolProperties,
		namesStringProperties,
		namesEntityIdProperties,
		namesIntProperties,
		namesArrayEntityIdProperties,
	)
	if err != nil {
		return nil, err
	}
	ed, err := entity.NewDataEntity(*id, *dataProperties)
	if err != nil {
		return nil, err
	}
	var libraryDataEntity entity.LibraryDataEntity = make(entity.LibraryDataEntity, 1)
	libraryDataEntity[nameDataEntity1] = *ed
	return &libraryDataEntity, nil

}
