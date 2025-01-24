package ValueModifierCommon

import "github.com/big-smiles/boardgame-golang/pkg/entity"

type ValueResolver[T any] interface {
	Resolve(
		executionVariables entity.Entity,
		managerPropertyId *entity.ManagerPropertyId,
	) (T, error)
}
