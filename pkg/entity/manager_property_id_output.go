package entity

type OutputManagerPropertyId struct {
	String        mapIdProperty[string]
	Int           mapIdProperty[int]
	Bool          mapIdProperty[bool]
	EntityId      mapIdProperty[Id]
	ArrayEntityId mapIdProperty[[]Id]
}

func NewOutputManagerPropertyId(
	string mapIdProperty[string],
	int mapIdProperty[int],
	bool mapIdProperty[bool],
	entityId mapIdProperty[Id],
	arrayEntityId mapIdProperty[[]Id],
) *OutputManagerPropertyId {
	return &OutputManagerPropertyId{
		String:        string,
		Int:           int,
		Bool:          bool,
		EntityId:      entityId,
		ArrayEntityId: arrayEntityId,
	}
}
