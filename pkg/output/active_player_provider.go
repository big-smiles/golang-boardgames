package output

import "github.com/big-smiles/golang-boardgames/pkg/player"

type ActivePlayerProvider interface {
	GetActivePlayers() ([]player.Id, error)
}
