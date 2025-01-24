package instruction

import "github.com/big-smiles/boardgame-golang/pkg/entity"

type ValueResolver struct {
	managerPropertyId *entity.ManagerPropertyId
}

func NewPerformerValueResolver() (*ValueResolver, error) {
	return &ValueResolver{}, nil
}

func (p *ValueResolver) Initialize(ctx InitializationContext) error {
	p.managerPropertyId = ctx.GetManagerPropertyId()
	return nil
}

func ResolveValueResolver[T any](
	executionVariables entity.Entity,
	performerValueResolver *ValueResolver,
	resolver IValueResolver[T],
) (T, error) {
	var zero T
	value, err := resolver.Resolve(executionVariables, performerValueResolver.managerPropertyId)
	if err != nil {
		return zero, err
	}
	return value, nil
}
