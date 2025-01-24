package entity

type Id int64
type ManagerEntityId struct {
	counter           Id
	idByName          map[NameEntityId]Id
	managerPropertyId *ManagerPropertyId
}

func NewManagerEntityId() (*ManagerEntityId, error) {
	m := make(map[NameEntityId]Id)
	return &ManagerEntityId{counter: 0,
		idByName: m}, nil
}

func (m *ManagerEntityId) GetId(
	executionVariables Entity,
	managerPropertyId *ManagerPropertyId,
	d DataId,
) (Id, error) {

	name, err := d.ResolverName.Resolve(executionVariables, managerPropertyId)
	if err != nil {
		return 0, err
	}
	if name == "" {
		return m.getNextId()
	}
	id, ok := m.idByName[name]
	if ok {
		return id, nil
	} else {
		id, err := m.getNextId()
		if err != nil {
			return 0, err
		}
		m.idByName[name] = id
		return id, nil
	}
}

func (m *ManagerEntityId) getNextId() (Id, error) {
	m.counter++
	return m.counter, nil
}

func (m *ManagerEntityId) Initialize(ctx IInitializationContext) error {
	m.managerPropertyId = ctx.GetManagerPropertyId()
	return nil
}
