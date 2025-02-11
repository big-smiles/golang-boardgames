package phase

type CallbackInstructionExecute func() error

type Stage struct {
	callback CallbackInstructionExecute
}

func NewStage(callback CallbackInstructionExecute) *Stage {
	return &Stage{callback: callback}
}
