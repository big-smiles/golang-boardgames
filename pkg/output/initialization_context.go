package output

import "github.com/big-smiles/boardgame-golang/pkg/entity"

type InitializationContext interface {
	GetManagerEntity() *entity.ManagerEntity
}
