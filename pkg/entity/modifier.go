package entity

type Modifier struct {
	propertiesModifier PropertiesModifier
}

func NewEntityModifier(executionVariables Entity, m *ManagerPropertyId, d DataModifier) (*Modifier, error) {
	p, err := NewPropertiesModifier(executionVariables, m, d.dataPropertiesModifier)
	if err != nil {
		return nil, err
	}
	return &Modifier{
		propertiesModifier: *p,
	}, nil
}
