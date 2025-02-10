package entity

type MapDataModifierProperties[T PropertyTypes] map[NamePropertyId[T]]IDataPropertyModifier[T]

type DataPropertiesModifier struct {
	IntModifiers           MapDataModifierProperties[int]
	StringModifiers        MapDataModifierProperties[string]
	BoolModifiers          MapDataModifierProperties[bool]
	EntityIdModifiers      MapDataModifierProperties[Id]
	ArrayEntityIdModifiers MapDataModifierProperties[[]Id]
}
