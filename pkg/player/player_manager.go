package player

type ManagerPlayer struct {
	Players []Id
}

func NewManagerPlayer(players []Id) (*ManagerPlayer, error) {
	return &ManagerPlayer{
		Players: players,
	}, nil
}

func (m *ManagerPlayer) Initialize(_ IInitializationContext) error {
	return nil
}
