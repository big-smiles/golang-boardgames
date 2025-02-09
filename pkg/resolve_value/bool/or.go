package resolveValueBool

import (
	"github.com/big-smiles/golang-boardgames/pkg/entity"
)

type ResolveValueOr struct {
	resolvers []BoolResolver
}

func NewResolveValueOr(resolvers ...BoolResolver) *ResolveValueOr {
	return &ResolveValueOr{resolvers: resolvers}
}
func (r ResolveValueOr) Resolve(
	variables entity.Entity,
	managerPropertyId *entity.ManagerPropertyId,
) (bool, error) {
	for _, resolver := range r.resolvers {
		result, err := resolver.Resolve(variables, managerPropertyId)
		if err != nil {
			return false, err
		}
		if result {
			return true, nil
		}
	}
	return false, nil
}
