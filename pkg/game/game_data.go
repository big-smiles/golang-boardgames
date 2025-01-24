package game

import (
	"github.com/big-smiles/golang-boardgames/pkg/entity"
	"github.com/big-smiles/golang-boardgames/pkg/phase"
	"github.com/big-smiles/golang-boardgames/pkg/player"
)

type DataGame struct {
	phases     phase.LibraryPhase
	firstPhase phase.NamePhase
	entities   entity.LibraryDataEntity
	players    []player.Id
}

func NewDataGame(
	entities entity.LibraryDataEntity,
	phases phase.LibraryPhase,
	firstPhase phase.NamePhase,
	players []player.Id,
) (*DataGame, error) {
	return &DataGame{
		phases:     phases,
		firstPhase: firstPhase,
		entities:   entities,
		players:    players,
	}, nil
}
