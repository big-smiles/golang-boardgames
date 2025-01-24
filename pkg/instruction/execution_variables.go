package instruction

import (
	"github.com/big-smiles/boardgame-golang/pkg/entity"
	resolveValueConstant "github.com/big-smiles/boardgame-golang/pkg/resolve_value/constant"
)

func getExecutionVariablesData() (*entity.DataEntity, error) {
	nameEntityIdResolver, err := resolveValueConstant.NewResolveConstant[entity.NameEntityId]("")
	if err != nil {
		return nil, err
	}
	dataId, err := entity.NewDataId(nameEntityIdResolver)
	if err != nil {
		return nil, err
	}
	dataProperties, err := entity.NewDataProperties(
		nil,
		nil,
		nil,
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return entity.NewDataEntity(
		*dataId,
		*dataProperties,
	)
}
