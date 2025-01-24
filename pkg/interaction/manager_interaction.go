package interaction

import (
	"fmt"
	"github.com/big-smiles/golang-boardgames/pkg/entity"
)

type Callback func([]OutputInteraction)
type Id uint64
type ManagerInteraction struct {
	availableInteractions      []AvailableInteraction
	idCounter                  Id
	waitingForInteraction      bool
	instructionTriggerer       ITriggerInstruction
	callbackOutputInteractions Callback
	managerPropertyId          *entity.ManagerPropertyId
}

func NewManagerInteraction(callbackOutputInteractions Callback) (*ManagerInteraction, error) {
	return &ManagerInteraction{
		availableInteractions:      make([]AvailableInteraction, 0),
		idCounter:                  0,
		callbackOutputInteractions: callbackOutputInteractions,
		waitingForInteraction:      false,
	}, nil
}

func (m *ManagerInteraction) Initialize(ctx IInitializationContext) error {
	m.instructionTriggerer = ctx.GetITriggerInstruction()
	m.managerPropertyId = ctx.GetManagerPropertyId()

	return nil
}

func (m *ManagerInteraction) AddAvailableInteraction(
	executionVariables entity.Entity,
	data DataAvailableInteraction,
	idToTrigger int,
) error {
	m.idCounter++
	ai, err := NewAvailableInteraction(
		executionVariables,
		m.managerPropertyId,
		m.idCounter,
		idToTrigger,
		data,
	)
	if err != nil {
		return err
	}
	m.availableInteractions = append(m.availableInteractions, *ai)
	return nil
}

func (m *ManagerInteraction) ClearAvailableInteraction() error {
	m.availableInteractions = make([]AvailableInteraction, 0)

	return nil
}

func (m *ManagerInteraction) SendOutputInteractions() (waitingForInteraction bool, err error) {
	if m.waitingForInteraction == false {
		return false, nil
	}
	if len(m.availableInteractions) == 0 {
		return false, nil
	}

	outputInteractions := make([]OutputInteraction, len(m.availableInteractions))
	for i, ai := range m.availableInteractions {
		a, err := NewOutputInteraction(ai)
		if err != nil {
			return false, err
		}
		outputInteractions[i] = *a
	}

	m.callbackOutputInteractions(outputInteractions)
	return true, nil
}

func (m *ManagerInteraction) ReceiveSelectedInteraction(interactions []SelectedInteraction) error {
	if !m.waitingForInteraction {
		return fmt.Errorf("manager_interaction not waiting for interaction")
	}
	m.waitingForInteraction = false
	selectedIndexes := make([]int, len(interactions))
	for k, interaction := range interactions {
		index, err := m.validateSelectedInteraction(interaction)
		if err != nil {
			return err
		}
		selectedIndexes[k] = index
	}
	for _, index := range selectedIndexes {
		err := m.instructionTriggerer.Trigger(
			m.availableInteractions[selectedIndexes[index]].instructionIdToTrigger,
			interactions[index].selectedEntities,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
func (m *ManagerInteraction) validateSelectedInteraction(interaction SelectedInteraction) (index int, err error) {
	for i, ai := range m.availableInteractions {
		if ai.Id == interaction.id {
			err := validateInteraction(interaction, ai)
			if err != nil {
				return i, err
			}

			return i, nil
		}
	}
	return 0, fmt.Errorf("available interaction id %d not found", interaction.id)
}
func (m *ManagerInteraction) WaitForInteraction() error {
	if m.waitingForInteraction {
		return fmt.Errorf("manager_interaction was already waiting for interaction")
	}
	m.waitingForInteraction = true
	return nil
}

func validateInteraction(selected SelectedInteraction, available AvailableInteraction) error {
	if selected.id != available.Id {
		return fmt.Errorf("interaction id %d is diferent from avaialable %d", selected.id, available.Id)
	}

	if len(selected.selectedEntities) < available.MinAmount {
		return fmt.Errorf("minimum amount of selectedEntities is %d", available.MinAmount)
	}

	if len(selected.selectedEntities) >= available.MaxAmount {
		return fmt.Errorf("maximum amount of selectedEntities is %d", available.MaxAmount)
	}

	mapEntities := make(map[entity.Id]bool, len(available.AvailableEntities))
	for _, ai := range available.AvailableEntities {
		mapEntities[ai] = true
	}
	for _, ai := range selected.selectedEntities {
		if !mapEntities[ai] {
			return fmt.Errorf("entity.Id %d was not found in available entities", ai)
		}
	}

	return nil
}
