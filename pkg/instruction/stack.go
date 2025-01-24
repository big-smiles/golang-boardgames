package instruction

type instructionStack []wrapperInstruction

func newInstructionStack() (*instructionStack, error) {
	var s []wrapperInstruction
	return (*instructionStack)(&s), nil
}

func (s *instructionStack) pop() (i wrapperInstruction, ok bool) {

	if len(*s) == 0 {
		return wrapperInstruction{}, false
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v, true
}

func (s *instructionStack) push(h wrapperInstruction) {
	*s = append(*s, h)
}
