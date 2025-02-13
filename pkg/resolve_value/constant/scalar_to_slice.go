package resolveValueConstant

import (
	"github.com/big-smiles/golang-boardgames/pkg/entity"
)

type ResolveScalarToSlice[T any] struct {
	value IValueResolver[T]
}

func NewResolveScalarToSlice[T any](value IValueResolver[T]) *ResolveScalarToSlice[T] {
	return &ResolveScalarToSlice[T]{
		value: value,
	}
}

func (r ResolveScalarToSlice[T]) Resolve(
	executionVariables entity.Entity,
	managerPropertyId *entity.ManagerPropertyId,
) ([]T, error) {
	var zero []T
	v, err := r.value.Resolve(executionVariables, managerPropertyId)
	if err != nil {
		return zero, err
	}
	return []T{v}, nil
}
