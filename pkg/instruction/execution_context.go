package instruction

import (
	"errors"
	"github.com/big-smiles/boardgame-golang/pkg/entity"
)

const (
	SELECTED_ENTITIES entity.NamePropertyId[[]entity.Id] = "__SELECTED_ENTITIES"
)

type ExecutionContext struct {
	Performer          *Performer
	ExecutionVariables entity.Entity
}

func newExecutionContext(performer *Performer, executionVariables entity.Entity) (*ExecutionContext, error) {
	if performer == nil {
		return nil, errors.New("performer must not be nil")
	}
	return &ExecutionContext{
		Performer:          performer,
		ExecutionVariables: executionVariables,
	}, nil
}