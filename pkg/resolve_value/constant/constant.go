package resolveValueConstant

import (
	"github.com/big-smiles/boardgame-golang/pkg/entity"
)

type ResolveValueConstant[T any] struct {
	value T
}

func NewResolveConstant[T any](v T) (*ResolveValueConstant[T], error) {
	return &ResolveValueConstant[T]{
		value: v,
	}, nil
}

func (r ResolveValueConstant[T]) Resolve(
	_ entity.Entity,
	_ *entity.ManagerPropertyId,
) (T, error) {
	return r.value, nil
}
