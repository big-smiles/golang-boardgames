package entity

type OutputEntity struct {
	Id         Id
	Name       NameEntityId
	Properties OutputProperties
}

func NewOutputEntity(entity Entity) (*OutputEntity, error) {
	var p OutputProperties
	if entity.properties != nil {
		aux, err := entity.properties.getOutput()
		if err != nil {
			return nil, err
		}
		p = *aux
	}

	return &OutputEntity{
		Id:         entity.Id,
		Name:       entity.Name,
		Properties: p,
	}, nil
}
