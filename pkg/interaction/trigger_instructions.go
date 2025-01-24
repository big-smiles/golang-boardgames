package interaction

import "github.com/big-smiles/golang-boardgames/pkg/entity"

type ITriggerInstruction interface {
	Trigger(idToTrigger int, selectedEntities []entity.Id) error
}
