package instruction

import (
	"github.com/big-smiles/golang-boardgames/pkg/phase"
)

type Phase struct {
	managerPhase *phase.ManagerPhase
}

func NewPerformerPhase() (*Phase, error) {
	return &Phase{}, nil
}
func (p *Phase) SetNextPhase(namePhase phase.NamePhase) error {
	err := p.managerPhase.SetNextPhase(namePhase)
	if err != nil {
		return err
	}
	return nil
}
func (p *Phase) Initialize(ctx InitializationContext) error {
	p.managerPhase = ctx.GetManagerPhase()
	return nil
}
