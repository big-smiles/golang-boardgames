package phase

import (
	"github.com/big-smiles/golang-boardgames/pkg/player"
)

type DataTurn struct {
	Name          NameTurn
	ActivePlayers []player.Id
	Stages        []DataStage
}
