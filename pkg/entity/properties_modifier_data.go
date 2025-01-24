package entity

type MapDataModifierProperties[T PropertyTypes] map[NamePropertyId[T]]IDataPropertyModifier[T]

type DataPropertiesModifier struct {
	intModifiers           MapDataModifierProperties[int]
	stringModifiers        MapDataModifierProperties[string]
	boolModifiers          MapDataModifierProperties[bool]
	entityIdModifiers      MapDataModifierProperties[Id]
	arrayEntityIdModifiers MapDataModifierProperties[[]Id]
}

func NewDataPropertiesModifier(
	intModifiers *MapDataModifierProperties[int],
	stringModifiers *MapDataModifierProperties[string],
	boolModifiers *MapDataModifierProperties[bool],
	entityIdModifiers *MapDataModifierProperties[Id],
	arrayEntityIdModifiers *MapDataModifierProperties[[]Id],
) (*DataPropertiesModifier, error) {
	if intModifiers == nil {
		intModifiers = &MapDataModifierProperties[int]{}
	}
	if stringModifiers == nil {
		stringModifiers = &MapDataModifierProperties[string]{}
	}
	if boolModifiers == nil {
		boolModifiers = &MapDataModifierProperties[bool]{}
	}
	if entityIdModifiers == nil {
		entityIdModifiers = &MapDataModifierProperties[Id]{}
	}
	if arrayEntityIdModifiers == nil {
		arrayEntityIdModifiers = &MapDataModifierProperties[[]Id]{}
	}
	return &DataPropertiesModifier{
		intModifiers:           *intModifiers,
		stringModifiers:        *stringModifiers,
		boolModifiers:          *boolModifiers,
		entityIdModifiers:      *entityIdModifiers,
		arrayEntityIdModifiers: *arrayEntityIdModifiers,
	}, nil
}
