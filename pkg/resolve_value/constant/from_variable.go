package resolveValueConstant

import (
	"github.com/big-smiles/golang-boardgames/pkg/entity"
)

type ResolveValueFromVariable[T entity.PropertyTypes] struct {
	namePropertyId entity.NamePropertyId[T]
}

func NewResolveValueFromVariable[T entity.PropertyTypes](namePropertyId entity.NamePropertyId[T]) (*ResolveValueFromVariable[T], error) {
	return &ResolveValueFromVariable[T]{
		namePropertyId: namePropertyId,
	}, nil
}

func (r ResolveValueFromVariable[T]) Resolve(
	executionVariables entity.Entity,
	managerPropertyId *entity.ManagerPropertyId,
) (T, error) {
	var zero T
	managerTypedPropertyId, err := entity.GetManagerTypedPropertyId[T](managerPropertyId)
	if err != nil {
		return zero, err
	}
	propertyId, err := managerTypedPropertyId.GetId(r.namePropertyId)
	if err != nil {
		return zero, err
	}
	value, err := entity.GetValueFromEntity(executionVariables, propertyId)
	if err != nil {
		return zero, err
	}
	return value, nil
}
