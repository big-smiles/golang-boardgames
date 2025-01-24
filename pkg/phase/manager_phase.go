package phase

import (
	"fmt"
	"github.com/big-smiles/boardgame-golang/pkg/entity"
	"github.com/big-smiles/boardgame-golang/pkg/instruction"
	"github.com/big-smiles/boardgame-golang/pkg/interaction"
)

type NamePhase string
type LibraryPhase map[NamePhase]DataPhase
type ManagerPhase struct {
	phases             LibraryPhase
	currentPhase       NamePhase
	indexCurrentTurn   int
	indexCurrentStage  int
	nextNamePhase      NamePhase
	nextPhaseSet       bool
	managerInstruction *instruction.ManagerInstruction
	managerInteraction *interaction.ManagerInteraction
}

func NewManagerPhase(
	phases map[NamePhase]DataPhase,
	firstPhase NamePhase,
) (*ManagerPhase, error) {
	return &ManagerPhase{
		phases:            phases,
		nextNamePhase:     firstPhase,
		currentPhase:      firstPhase,
		indexCurrentTurn:  0,
		indexCurrentStage: 0,
		nextPhaseSet:      false,
	}, nil
}
func (m *ManagerPhase) Initialize(ctx IInitializationContext) error {
	m.managerInstruction = ctx.GetManagerInstruction()
	m.managerInteraction = ctx.GetManagerInteraction()
	if m.managerInstruction == nil {
		return fmt.Errorf("no ManagerInstruction")
	}
	if m.managerInteraction == nil {
		return fmt.Errorf("no managerInteraction")
	}
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
	if m.indexCurrentTurn >= len(m.phases[m.currentPhase].turns) {
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
	if m.indexCurrentStage >= len(m.phases[m.currentPhase].turns[m.indexCurrentTurn].stages) {
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
	if m.indexCurrentTurn >= len(m.phases[m.currentPhase].turns) {
		return fmt.Errorf(
			"invalid turn index= %d for phase with turn length of=%d",
			m.indexCurrentTurn,
			len(m.phases[m.currentPhase].turns),
		)
	}
	if m.indexCurrentStage >= len(m.phases[m.currentPhase].turns[m.indexCurrentTurn].stages) {
		return fmt.Errorf(
			"invalid stage index= %d for turn with stage length of=%d",
			m.indexCurrentStage,
			len(m.phases[m.currentPhase].turns[m.indexCurrentTurn].stages),
		)
	}
	i, err := m.phases[m.currentPhase].turns[m.indexCurrentTurn].stages[m.indexCurrentStage].instructions.NewFromThisData()
	if err != nil {
		return err
	}
	err = m.managerInstruction.AddInstruction(i, []entity.Id{})
	if err != nil {
		return err
	}
	return nil
}
func (m *ManagerPhase) SetNextPhase(namePhase NamePhase) error {
	m.nextPhaseSet = true
	_, ok := m.phases[namePhase]
	if !ok {
		return fmt.Errorf("invalid phase name: %s", namePhase)
	}
	m.nextNamePhase = namePhase
	return nil
}
