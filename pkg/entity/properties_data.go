package entity

type DataPropertiesTyped[T PropertyTypes] []NamePropertyId[T]
type DataProperties struct {
	BoolProperties          DataPropertiesTyped[bool]
	StringProperties        DataPropertiesTyped[string]
	EntityIdProperties      DataPropertiesTyped[Id]
	IntProperties           DataPropertiesTyped[int]
	ArrayEntityIdProperties DataPropertiesTyped[[]Id]
}
