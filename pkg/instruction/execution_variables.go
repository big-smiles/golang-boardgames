package instruction

import (
	"github.com/big-smiles/golang-boardgames/pkg/entity"
	resolveValueConstant "github.com/big-smiles/golang-boardgames/pkg/resolve_value/constant"
)

func getExecutionVariablesData() (*entity.DataEntity, error) {
	nameEntityIdResolver := resolveValueConstant.NewResolveConstant[entity.NameEntityId]("")
	dataId, err := entity.NewDataId(nameEntityIdResolver)
	if err != nil {
		return nil, err
	}

	dataProperties := entity.DataProperties{}

	return entity.NewDataEntity(
		*dataId,
		dataProperties,
	)
}
