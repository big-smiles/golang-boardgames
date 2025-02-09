package resolveValueBool

import (
	"github.com/big-smiles/golang-boardgames/pkg/entity"
)

type ResolveValueEquals[T comparable] struct {
	resolvers []IValueResolver[T]
}

func NewResolveValueEquals[T comparable](resolvers ...IValueResolver[T]) *ResolveValueEquals[T] {
	return &ResolveValueEquals[T]{
		resolvers: resolvers,
	}
}
func (r ResolveValueEquals[T]) Resolve(
	variables entity.Entity,
	managerPropertyId *entity.ManagerPropertyId,
) (bool, error) {
	first := true
	var value T
	for _, resolver := range r.resolvers {
		result, err := resolver.Resolve(variables, managerPropertyId)
		if err != nil {
			return false, err
		}
		if first {
			first = false
			value = result
		} else {
			if value != result {
				return false, nil
			}
		}
	}
	return true, nil
}
