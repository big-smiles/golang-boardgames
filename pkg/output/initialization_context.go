package output

import "github.com/big-smiles/golang-boardgames/pkg/entity"

type InitializationContext interface {
	GetManagerEntity() *entity.ManagerEntity
}
