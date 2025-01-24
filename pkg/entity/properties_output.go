package entity

type mapOutputProperties[T PropertyTypes] map[PropertyId[T]]T
type OutputProperties struct {
	IntProperties           mapOutputProperties[int]
	StringProperties        mapOutputProperties[string]
	BoolProperties          mapOutputProperties[bool]
	EntityIdProperties      mapOutputProperties[Id]
	ArrayEntityIdProperties mapOutputProperties[[]Id]
}

func NewOutputProperties(
	intProperties *mapOutputProperties[int],
	stringProperties *mapOutputProperties[string],
	boolProperties *mapOutputProperties[bool],
	entityIdProperties *mapOutputProperties[Id],
	arrayEntityIdProperties *mapOutputProperties[[]Id],
) (*OutputProperties, error) {
	return &OutputProperties{
		IntProperties:           *intProperties,
		StringProperties:        *stringProperties,
		BoolProperties:          *boolProperties,
		EntityIdProperties:      *entityIdProperties,
		ArrayEntityIdProperties: *arrayEntityIdProperties,
	}, nil
}
