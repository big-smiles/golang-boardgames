package resolveValueConstant

import (
	"github.com/big-smiles/golang-boardgames/pkg/entity"
)

type ResolveValueConstant[T any] struct {
	value T
}

func NewResolveConstant[T any](v T) *ResolveValueConstant[T] {
	return &ResolveValueConstant[T]{
		value: v,
	}
}

func (r ResolveValueConstant[T]) Resolve(
	_ entity.Entity,
	_ *entity.ManagerPropertyId,
) (T, error) {
	return r.value, nil
}
