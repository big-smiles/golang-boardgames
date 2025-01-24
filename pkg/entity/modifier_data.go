package entity

type DataModifier struct {
	dataPropertiesModifier DataPropertiesModifier
}

func NewDataEntityModifier(d DataPropertiesModifier) (*DataModifier, error) {
	return &DataModifier{
		dataPropertiesModifier: d,
	}, nil
}
