package phase

import (
	"github.com/big-smiles/golang-boardgames/pkg/player"
	"testing"
)

type MockInitializationContext struct {
}

const (
	firstPhase  NamePhase = "first_phase"
	secondPhase NamePhase = "second_phase"
	player1     player.Id = "player1"
	player2     player.Id = "player2"
)

func TestManagerPhase(t *testing.T) {
	managerPhase, err := NewManagerPhase()
	if err != nil {
		t.Fatal(err)
	}
	ctx := MockInitializationContext{}
	err = managerPhase.Initialize(ctx)
	if err != nil {
		t.Fatal(err)
	}
	callbackStage1Called := 0
	callbackStage2Called := 0
	callbackStage3Called := 0
	callbackStage1 := func() error {
		callbackStage1Called++
		err = managerPhase.SetNextPhase(secondPhase)
		if err != nil {
			return err
		}
		return nil
	}
	callbackStage2 := func() error {
		callbackStage2Called++
		return nil
	}
	callbackStage3 := func() error {
		callbackStage3Called++
		return nil
	}
	libraryPhase, firstPhase := getPhaseData(callbackStage1, callbackStage2, callbackStage3)
	err = managerPhase.LoadPhases(libraryPhase, firstPhase)
	if err != nil {
		t.Fatal(err)
	}

	err = managerPhase.Next()
	if err != nil {
		t.Fatal(err)
	}
	if callbackStage1Called != 1 || callbackStage2Called != 0 || callbackStage3Called != 0 {
		t.Error("failed after first Next", callbackStage1Called, callbackStage2Called)
	}

	err = managerPhase.Next()
	if err != nil {
		t.Fatal(err)
	}
	if callbackStage1Called != 1 || callbackStage2Called != 1 || callbackStage3Called != 0 {
		t.Error("failed after second Next", callbackStage1Called, callbackStage2Called)
	}
	activePlayers, err := managerPhase.GetActivePlayers()
	if err != nil {
		t.Fatal(err)
	}
	if len(activePlayers) != 1 || activePlayers[0] != player1 {
		t.Error("failed after GetActivePlayers", activePlayers)
	}

	err = managerPhase.Next()
	if err != nil {
		t.Fatal(err)
	}
	if callbackStage1Called != 1 || callbackStage2Called != 1 || callbackStage3Called != 1 {
		t.Error("failed after second Next", callbackStage1Called, callbackStage2Called)
	}
	activePlayers, err = managerPhase.GetActivePlayers()
	if err != nil {
		t.Fatal(err)
	}
	if len(activePlayers) != 1 || activePlayers[0] != player2 {
		t.Error("failed after GetActivePlayers", activePlayers)
	}
}

func getPhaseData(
	callbackStage1 CallbackInstructionExecute,
	callbackStage2 CallbackInstructionExecute,
	callbackStage3 CallbackInstructionExecute,
) (LibraryPhase, NamePhase) {

	return LibraryPhase{
		firstPhase: {
			Name: firstPhase,
			Turns: []Turn{
				{
					Stages: []Stage{
						{
							callback: callbackStage1,
						},
					},
				},
			},
		},
		secondPhase: {
			Name: secondPhase,
			Turns: []Turn{
				{
					ActivePlayers: []player.Id{player1},
					Stages: []Stage{
						{
							callback: callbackStage2,
						},
					},
				},
				{
					ActivePlayers: []player.Id{player2},
					Stages: []Stage{
						{
							callback: callbackStage3,
						},
					},
				},
			},
		},
	}, firstPhase
}
