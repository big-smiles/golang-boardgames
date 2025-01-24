package entity

type NameEntityId string
type DataId struct {
	ResolverName IValueResolver[NameEntityId]
}

func NewDataId(resolverName IValueResolver[NameEntityId]) (*DataId, error) {
	return &DataId{
		ResolverName: resolverName,
	}, nil
}
