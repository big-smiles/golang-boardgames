package instruction

type Instruction interface {
	Execute(ctx ExecutionContext) error
}
type DataInstruction interface {
	NewFromThisData() (Instruction, error)
}
