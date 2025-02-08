package resolveValueBool

import "github.com/big-smiles/golang-boardgames/pkg/entity"

type BoolResolver interface {
	Resolve(
		executionVariables entity.Entity,
		managerPropertyId *entity.ManagerPropertyId,
	) (bool, error)
}
