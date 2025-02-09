package instruction

import (
	"errors"
	"github.com/big-smiles/golang-boardgames/pkg/output"
)

type Output struct {
	managerOutput *output.ManagerOutput
}

func NewPerformerOutput() (*Output, error) {
	return &Output{}, nil
}

func (p *Output) SendOutput() error {
	if p.managerOutput == nil {
		return errors.New("output does not have a managerOutput")
	}
	err := p.managerOutput.SendOutput()
	if err != nil {
		return err
	}
	return nil

}
func (p *Output) Initialize(ctx InitializationContext) error {
	p.managerOutput = ctx.GetManagerOutput()
	return nil
}
