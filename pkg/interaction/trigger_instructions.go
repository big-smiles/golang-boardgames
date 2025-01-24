package interaction

import "github.com/big-smiles/boardgame-golang/pkg/entity"

type ITriggerInstruction interface {
	Trigger(idToTrigger int, selectedEntities []entity.Id) error
}
