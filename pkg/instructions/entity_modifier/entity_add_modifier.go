package instructionEntityModifier

import (
	"github.com/big-smiles/boardgame-golang/pkg/entity"
	"github.com/big-smiles/boardgame-golang/pkg/instruction"
)

type instructionAddEntityModifier struct {
	target             entity.DataId
	dataEntityModifier entity.DataModifier
}

func (i instructionAddEntityModifier) Execute(ctx instruction.ExecutionContext) error {
	id, err := ctx.Performer.Entity.GetId(ctx.ExecutionVariables, i.target)
	if err != nil {
		return err
	}

	err = ctx.Performer.Entity.AddModifier(ctx.ExecutionVariables, []entity.Id{id}, i.dataEntityModifier)
	if err != nil {
		return err
	}
	return nil
}

func newInstructionAddEntityModifier(d DataInstructionAddEntityModifier) (*instructionAddEntityModifier, error) {
	a := &instructionAddEntityModifier{
		target:             d.target,
		dataEntityModifier: d.dataEntityModifier,
	}
	return a, nil
}
