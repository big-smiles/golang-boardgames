package resolveValueBool

import (
	"github.com/big-smiles/golang-boardgames/pkg/entity"
)

type ResolveValueAnd struct {
	resolvers []BoolResolver
}

func NewResolveValueAnd(resolvers ...BoolResolver) *ResolveValueAnd {
	return &ResolveValueAnd{resolvers}
}
func (r ResolveValueAnd) Resolve(
	variables entity.Entity,
	managerPropertyId *entity.ManagerPropertyId,
) (bool, error) {
	for _, resolver := range r.resolvers {
		result, err := resolver.Resolve(variables, managerPropertyId)
		if err != nil {
			return false, err
		}
		if !result {
			return false, nil
		}
	}
	return true, nil
}
