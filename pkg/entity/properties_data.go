package entity

type DataPropertiesTyped[T PropertyTypes] []NamePropertyId[T]
type DataProperties struct {
	boolProperties          DataPropertiesTyped[bool]
	stringProperties        DataPropertiesTyped[string]
	entityIdProperties      DataPropertiesTyped[Id]
	intProperties           DataPropertiesTyped[int]
	arrayEntityIdProperties DataPropertiesTyped[[]Id]
}

func NewDataProperties(
	boolProperties DataPropertiesTyped[bool],
	stringProperties DataPropertiesTyped[string],
	entityIdProperties DataPropertiesTyped[Id],
	intProperties DataPropertiesTyped[int],
	arrayEntityIdProperties DataPropertiesTyped[[]Id],
) (*DataProperties, error) {
	return &DataProperties{
		boolProperties:          boolProperties,
		stringProperties:        stringProperties,
		entityIdProperties:      entityIdProperties,
		intProperties:           intProperties,
		arrayEntityIdProperties: arrayEntityIdProperties,
	}, nil
}
