package output

import (
	"errors"
	"github.com/big-smiles/boardgame-golang/pkg/entity"
	"github.com/big-smiles/boardgame-golang/pkg/interaction"
)

type Callback func(output *Game)

type ManagerOutput struct {
	managerEntity      *entity.ManagerEntity
	managerInteraction *interaction.ManagerInteraction
	callback           func(output *Game)
}

func NewManagerOutput(c func(output *Game)) (*ManagerOutput, error) {
	return &ManagerOutput{
		callback: c,
	}, nil
}
func (m *ManagerOutput) Initialize(ctx InitializationContext) (err error) {
	m.managerEntity = ctx.GetManagerEntity()
	if m.managerEntity == nil {
		return errors.New("managerEntity is nil")
	}
	return nil
}

func (m *ManagerOutput) SendOutput() error {
	amount := m.managerEntity.GetOutputAmount()
	o := make([]entity.OutputEntity, amount)
	err := m.managerEntity.GetOutput(&o)
	if err != nil {

		return err
	}

	eo, err := NewGameOutput(o)
	if err != nil {
		return err
	}

	m.callback(eo)
	return nil
}
