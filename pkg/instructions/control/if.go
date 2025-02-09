package instruction_control

import "github.com/big-smiles/golang-boardgames/pkg/instruction"

type InstructionIf struct {
	condition IValueResolver[bool]
	whenTrue  instruction.DataInstruction
	whenFalse instruction.DataInstruction
}

func newInstructionIf(
	data DataInstructionIf,
) (*InstructionIf, error) {
	return &InstructionIf{
		condition: data.condition,
		whenTrue:  data.whenTrue,
		whenFalse: data.whenFalse,
	}, nil
}

func (i InstructionIf) Execute(ctx instruction.ExecutionContext) error {
	result, err := instruction.ResolveValueResolver[bool](
		ctx.ExecutionVariables,
		ctx.Performer.ValueResolver,
		i.condition,
	)
	if err != nil {
		return err
	}
	if result {
		whenTrue, err := i.whenTrue.NewFromThisData()
		if err != nil {
			return err
		}
		return whenTrue.Execute(ctx)
	} else {
		whenFalse, err := i.whenFalse.NewFromThisData()
		if err != nil {
			return err
		}
		return whenFalse.Execute(ctx)
	}
}
