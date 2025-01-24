package phase

import (
	"github.com/big-smiles/golang-boardgames/pkg/player"
)

type DataTurn struct {
	activePlayers []player.Id
	stages        []DataStage
}

func NewDataTurn(players []player.Id, stages []DataStage) (*DataTurn, error) {
	return &DataTurn{
		activePlayers: players,
		stages:        stages,
	}, nil
}
