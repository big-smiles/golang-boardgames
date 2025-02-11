package phaseData

import "github.com/big-smiles/golang-boardgames/pkg/phase"

type LibraryDataPhase map[phase.NamePhase]DataPhase

type DataPhase struct {
	Name  phase.NamePhase
	Turns []DataTurn
}
