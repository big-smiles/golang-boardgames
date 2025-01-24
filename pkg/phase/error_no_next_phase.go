package phase

type ErrorNoNextPhase struct {
	message string
}

func (e ErrorNoNextPhase) Error() string {
	return e.message
}

func NewErrorNoNextPhase() ErrorNoNextPhase {
	return ErrorNoNextPhase{message: "no next phase was set"}
}
