package resolveValueBool

import (
	"github.com/big-smiles/golang-boardgames/pkg/entity"
)

type ResolveValueAnd struct {
	Resolvers []BoolResolver
}

func (r ResolveValueAnd) Resolve(
	variables entity.Entity,
	managerPropertyId *entity.ManagerPropertyId,
) (bool, error) {
	for _, resolver := range r.Resolvers {
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
