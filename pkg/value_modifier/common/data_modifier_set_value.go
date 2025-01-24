package ValueModifierCommon

import "github.com/big-smiles/boardgame-golang/pkg/entity"

type DataModifierSetValue[T entity.PropertyTypes] struct {
	value ValueResolver[T]
}

func NewDataModifierSetValue[T entity.PropertyTypes](value ValueResolver[T]) (*DataModifierSetValue[T], error) {
	return &DataModifierSetValue[T]{
		value: value,
	}, nil
}
func (d *DataModifierSetValue[T]) NewFromThisData(executionVariables entity.Entity, managerPropertyId *entity.ManagerPropertyId) (entity.IPropertyModifier[T], error) {
	v, err := d.value.Resolve(executionVariables, managerPropertyId)
	if err != nil {
		return nil, err
	}
	m, err := NewModifierSetValue[T](v)
	if err != nil {
		return nil, err
	}
	return m, nil
}
