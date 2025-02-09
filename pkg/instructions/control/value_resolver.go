package instruction_control

import "github.com/big-smiles/golang-boardgames/pkg/entity"

type IValueResolver[T any] interface {
	Resolve(
		executionVariables entity.Entity,
		managerPropertyId *entity.ManagerPropertyId,
	) (T, error)
}
