package ValueModifierCommon

import "github.com/big-smiles/golang-boardgames/pkg/entity"

type ValueResolver[T any] interface {
	Resolve(
		executionVariables entity.Entity,
		managerPropertyId *entity.ManagerPropertyId,
	) (T, error)
}
