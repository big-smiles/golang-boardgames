package phase

import "github.com/big-smiles/golang-boardgames/pkg/player"

type Turn struct {
	Name          NameTurn
	ActivePlayers []player.Id
	Stages        []Stage
}

func NewTurn(
	Name NameTurn,
	ActivePlayers []player.Id,
	Stages []Stage,
) *Turn {
	return &Turn{
		Name:          Name,
		ActivePlayers: ActivePlayers,
		Stages:        Stages,
	}
}
