package phase

type DataPhase struct {
	turns []DataTurn
}

func NewDataPhase(turns []DataTurn) (*DataPhase, error) {
	return &DataPhase{
		turns: turns,
	}, nil
}
