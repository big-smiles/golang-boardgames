package instruction_phase

import (
	"github.com/big-smiles/golang-boardgames/pkg/instruction"
	"github.com/big-smiles/golang-boardgames/pkg/phase"
)

type InstructionSetNextPhase struct {
	phaseName IValueResolver[phase.NamePhase]
}

func (i InstructionSetNextPhase) Execute(ctx instruction.ExecutionContext) error {
	resolved, err := instruction.ResolveValueResolver(ctx.ExecutionVariables, ctx.Performer.ValueResolver, i.phaseName)
	if err != nil {
		return err
	}
	err = ctx.Performer.Phase.SetNextPhase(resolved)
	if err != nil {
		return err
	}
	return nil
}

func newInstructionSetNextPhase(d DataInstructionSetNextPhase) (*InstructionSetNextPhase, error) {
	return &InstructionSetNextPhase{
		phaseName: d.phaseName,
	}, nil
}
