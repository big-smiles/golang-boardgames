package instruction

import (
	"errors"
	"fmt"
	"github.com/big-smiles/boardgame-golang/pkg/entity"
)

type ManagerTriggerInstruction struct {
	instructionsToTrigger map[int]DataInstruction
	managerInstruction    *ManagerInstruction
	counter               int
}

func NewManagerTriggerInstruction() (*ManagerTriggerInstruction, error) {
	return &ManagerTriggerInstruction{
		instructionsToTrigger: map[int]DataInstruction{},
		counter:               0,
	}, nil
}

func (i *ManagerTriggerInstruction) Initialize(ctx InitializationContext) error {
	i.managerInstruction = ctx.GetManagerInstruction()
	return nil
}

func (i *ManagerTriggerInstruction) Trigger(idToTrigger int, selectedEntities []entity.Id) error {
	if i.managerInstruction == nil {
		return errors.New("managerInstruction is nil on ManagerTriggerInstruction")
	}
	data, ok := i.instructionsToTrigger[idToTrigger]
	if !ok {
		return fmt.Errorf("no instruction to trigger with idToTrigger=%d", idToTrigger)
	}
	instruction, err := data.NewFromThisData()
	if err != nil {
		return err
	}
	err = i.managerInstruction.AddInstruction(instruction, selectedEntities)
	if err != nil {
		return err
	}

	return nil
}
func (i *ManagerTriggerInstruction) AddInstructionToTrigger(dataInstruction DataInstruction) (int, error) {
	i.counter++
	i.instructionsToTrigger[i.counter] = dataInstruction
	return i.counter, nil
}
func (i *ManagerTriggerInstruction) RemoveInstructionToTrigger(id int) error {
	delete(i.instructionsToTrigger, id)
	return nil
}
