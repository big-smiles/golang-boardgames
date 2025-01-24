package instructionEntityModifier

import "github.com/big-smiles/boardgame-golang/pkg/entity"

type IValueResolver[T any] interface {
	Resolve(
		executionVariables entity.Entity,
		managerPropertyId *entity.ManagerPropertyId,
	) (T, error)
}
