package entity

type DataEntity struct {
	id             DataId
	dataProperties DataProperties
}

func NewDataEntity(id DataId, dataProperties DataProperties) (*DataEntity, error) {
	return &DataEntity{
		id:             id,
		dataProperties: dataProperties,
	}, nil
}
