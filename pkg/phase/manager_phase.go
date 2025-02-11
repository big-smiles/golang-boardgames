package phase

import (
	"fmt"
	"github.com/big-smiles/golang-boardgames/pkg/interaction"
	"github.com/big-smiles/golang-boardgames/pkg/player"
)

type NamePhase string
type NameTurn string
type LibraryPhase map[NamePhase]Phase
type ManagerPhase struct {
	phases             LibraryPhase
	currentPhase       NamePhase
	indexCurrentTurn   int
	indexCurrentStage  int
	nextNamePhase      NamePhase
	nextPhaseSet       bool
	managerInteraction *interaction.ManagerInteraction
}

func NewManagerPhase() (*ManagerPhase, error) {
	return &ManagerPhase{
		indexCurrentTurn:  0,
		indexCurrentStage: 0,
		nextPhaseSet:      false,
	}, nil
}
func (m *ManagerPhase) Initialize(ctx IInitializationContext) error {
	m.managerInteraction = ctx.GetManagerInteraction()
	if m.managerInteraction == nil {
		return fmt.Errorf("no managerInteraction")
	}
	return nil
}
func (m *ManagerPhase) LoadPhases(
	phases LibraryPhase,
	firstPhase NamePhase) error {
	m.phases = phases
	m.currentPhase = firstPhase
	m.nextNamePhase = firstPhase
	return nil
}
func (m *ManagerPhase) Next() error {
	next, err := m.nextStage()
	if err != nil {
		return err
	}
	if !next {
		return nil
	}

	next, err = m.nextTurn()
	if err != nil {
		return err
	}
	if !next {
		return nil
	}

	err = m.nextPhase()
	if err != nil {
		return err
	}

	return nil
}
func (m *ManagerPhase) nextPhase() error {
	if !m.nextPhaseSet {
		return NewErrorNoNextPhase()
	}
	m.nextPhaseSet = false
	m.currentPhase = m.nextNamePhase
	m.indexCurrentTurn = 0
	m.indexCurrentStage = 0
	err := m.runCurrentStageInstruction()
	if err != nil {
		return err
	}
	return nil
}
func (m *ManagerPhase) nextTurn() (nextPhase bool, err error) {
	m.indexCurrentTurn++
	if m.indexCurrentTurn >= len(m.phases[m.currentPhase].Turns) {
		return true, nil
	}
	m.indexCurrentStage = 0
	err = m.runCurrentStageInstruction()
	if err != nil {
		return false, err
	}
	return false, nil
}
func (m *ManagerPhase) nextStage() (nextTurn bool, err error) {
	if m.indexCurrentStage >= len(m.phases[m.currentPhase].Turns[m.indexCurrentTurn].Stages) {
		return true, nil
	}
	err = m.runCurrentStageInstruction()
	if err != nil {
		return false, err
	}
	m.indexCurrentStage++
	return false, nil
}
func (m *ManagerPhase) runCurrentStageInstruction() error {
	if _, ok := m.phases[m.currentPhase]; !ok {
		return fmt.Errorf("current phase not on the phases library phaseName=%s", m.currentPhase)
	}
	if m.indexCurrentTurn >= len(m.phases[m.currentPhase].Turns) {
		return fmt.Errorf(
			"invalid turn index= %d for phase with turn length of=%d",
			m.indexCurrentTurn,
			len(m.phases[m.currentPhase].Turns),
		)
	}
	if m.indexCurrentStage >= len(m.phases[m.currentPhase].Turns[m.indexCurrentTurn].Stages) {
		return fmt.Errorf(
			"invalid stage index= %d for turn with stage length of=%d",
			m.indexCurrentStage,
			len(m.phases[m.currentPhase].Turns[m.indexCurrentTurn].Stages),
		)
	}
	err := m.phases[m.currentPhase].Turns[m.indexCurrentTurn].Stages[m.indexCurrentStage].callback()
	if err != nil {
		return err
	}
	return nil
}

func (m *ManagerPhase) SetNextPhase(namePhase NamePhase) error {
	m.nextPhaseSet = true
	_, ok := m.phases[namePhase]
	if !ok {
		return fmt.Errorf("invalid phase Name: %s", namePhase)
	}
	m.nextNamePhase = namePhase
	return nil
}

func (m *ManagerPhase) GetActivePlayers() ([]player.Id, error) {
	phaseData := m.phases[m.currentPhase]
	turnData := phaseData.Turns[m.indexCurrentTurn]
	return turnData.ActivePlayers, nil
}
