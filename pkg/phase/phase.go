package phase

type Phase struct {
	Name  NamePhase
	Turns []Turn
}

func NewPhase(name NamePhase, turns []Turn) *Phase {
	return &Phase{
		Name:  name,
		Turns: turns,
	}
}
