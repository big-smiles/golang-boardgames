package output

import (
	"errors"
	"github.com/big-smiles/golang-boardgames/pkg/entity"
)

type Callback func(output *Game)

type ManagerOutput struct {
	managerEntity     *entity.ManagerEntity
	managerPropertyId *entity.ManagerPropertyId
	callback          func(output *Game)
}

func NewManagerOutput(c func(output *Game)) (*ManagerOutput, error) {
	return &ManagerOutput{
		callback: c,
	}, nil
}
func (m *ManagerOutput) Initialize(ctx InitializationContext) (err error) {
	m.managerEntity = ctx.GetManagerEntity()
	m.managerPropertyId = ctx.GetManagerPropertyId()
	if m.managerEntity == nil {
		return errors.New("managerEntity is nil")
	}
	return nil
}

func (m *ManagerOutput) SendOutput() error {
	amount := m.managerEntity.GetOutputAmount()
	entities := make([]entity.OutputEntity, amount)
	err := m.managerEntity.GetOutput(&entities)
	if err != nil {

		return err
	}
	propertyIds, err := m.managerPropertyId.GetOutput()
	if err != nil {
		return err
	}
	eo, err := NewGameOutput(entities, propertyIds)
	if err != nil {
		return err
	}

	m.callback(eo)
	return nil
}
