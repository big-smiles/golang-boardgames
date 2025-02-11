package phaseData

import (
	"github.com/big-smiles/golang-boardgames/pkg/phase"
	"github.com/big-smiles/golang-boardgames/pkg/player"
)

type DataTurn struct {
	Name          phase.NameTurn
	ActivePlayers []player.Id
	Stages        []DataStage
}
