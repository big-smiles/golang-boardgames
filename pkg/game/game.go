package game

import (
	"fmt"
	"github.com/big-smiles/golang-boardgames/pkg/engine"
	"github.com/big-smiles/golang-boardgames/pkg/interaction"
	"github.com/big-smiles/golang-boardgames/pkg/output"
)

type ManagerGame struct {
	engineContext engine.EngineContext
	started       bool
}

func NewGame(
	data DataGame,
	outputCallback output.Callback,
	interactionCallback interaction.Callback,
) (*ManagerGame, error) {
	ctx, err := engine.NewEngineContext(
		data.entities,
		outputCallback,
		data.phases,
		data.firstPhase,
		interactionCallback,
		data.players,
	)
	if err != nil {
		return nil, err
	}

	return &ManagerGame{
		engineContext: *ctx,
		started:       false,
	}, nil
}
func (g *ManagerGame) Start() error {
	if g.started == true {
		return fmt.Errorf("engine already started")
	}
	g.started = true
	err := g.gameLoop()
	if err != nil {
		return err
	}
	return nil
}
func (g *ManagerGame) SelectInteraction(selectedInteractions []interaction.SelectedInteraction) error {
	err := g.engineContext.ManagerInteraction.ReceiveSelectedInteraction(selectedInteractions)
	if err != nil {
		return err
	}
	waitingForInteraction, err := g.engineContext.ManagerInteraction.SendOutputInteractions()
	if err != nil {
		return err
	}

	if waitingForInteraction == false {
		return g.gameLoop()
	}
	return nil
}
func (g *ManagerGame) gameLoop() error {
	err := g.engineContext.Next()
	if err != nil {
		return err
	}

	waitingForInteraction, err := g.engineContext.ManagerInteraction.SendOutputInteractions()
	if err != nil {
		return err
	}

	if waitingForInteraction == false {
		return g.gameLoop()
	}
	return nil
}
