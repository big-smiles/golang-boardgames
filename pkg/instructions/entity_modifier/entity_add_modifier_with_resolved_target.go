package instructionEntityModifier

import (
	"github.com/big-smiles/boardgame-golang/pkg/entity"
	"github.com/big-smiles/boardgame-golang/pkg/instruction"
)

type instructionAddEntityModifierWithResolvedTarget struct {
	target             IValueResolver[[]entity.Id]
	dataEntityModifier entity.DataModifier
}

func (i instructionAddEntityModifierWithResolvedTarget) Execute(ctx instruction.ExecutionContext) error {
	targetId, err := instruction.ResolveValueResolver(ctx.ExecutionVariables, ctx.Performer.ValueResolver, i.target)
	if err != nil {
		return err
	}

	err = ctx.Performer.Entity.AddModifier(ctx.ExecutionVariables, targetId, i.dataEntityModifier)
	if err != nil {
		return err
	}
	return nil
}
func newInstructionAddEntityModifierWithResolvedTarget(
	d DataInstructionAddEntityModifierWithResolvedTarget,
) (*instructionAddEntityModifierWithResolvedTarget, error) {
	a := &instructionAddEntityModifierWithResolvedTarget{
		target:             d.target,
		dataEntityModifier: d.dataEntityModifier,
	}
	return a, nil
}
