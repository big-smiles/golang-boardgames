package instruction_entity

import (
	"fmt"
	"github.com/big-smiles/boardgame-golang/pkg/entity"
	"github.com/big-smiles/boardgame-golang/pkg/instruction"
	resolveValueConstant "github.com/big-smiles/boardgame-golang/pkg/resolve_value/constant"
	ValueModifierCommon "github.com/big-smiles/boardgame-golang/pkg/value_modifier/common"
)

type instructionCreateEntityIntoVariable struct {
	nameDataEntity       entity.NameDataEntity
	variablePropertyName entity.NamePropertyId[entity.Id]
}

func (i instructionCreateEntityIntoVariable) Execute(ctx instruction.ExecutionContext) error {
	data, err := ctx.Performer.Entity.GetData(i.nameDataEntity)
	if err != nil {
		return err
	}
	id, err := ctx.Performer.Entity.Create(ctx.ExecutionVariables, data)
	fmt.Printf("created entity id=%d", id)
	if err != nil {
		return err
	}
	valueResolver, err := resolveValueConstant.NewResolveConstant[entity.Id](id)
	if err != nil {
		return err
	}
	setValueModifier, err := ValueModifierCommon.NewDataModifierSetValue(valueResolver)
	if err != nil {
		return err
	}
	entityIdMapDataModifierProperties := make(entity.MapDataModifierProperties[entity.Id], 1)
	entityIdMapDataModifierProperties[i.variablePropertyName] = setValueModifier

	dataPropertiesModifier, err := entity.NewDataPropertiesModifier(
		nil,
		nil,
		nil,
		&entityIdMapDataModifierProperties,
		nil,
	)
	if err != nil {
		return err
	}

	dataModifier, err := entity.NewDataEntityModifier(*dataPropertiesModifier)
	if err != nil {
		return err
	}

	err = ctx.Performer.Entity.AddModifier(ctx.ExecutionVariables, []entity.Id{ctx.ExecutionVariables.Id}, *dataModifier)

	return nil
}

func newInstructionCreateEntityIntoVariable(d DataInstructionCreateEntityIntoVariable) (*instructionCreateEntityIntoVariable, error) {
	a := &instructionCreateEntityIntoVariable{
		nameDataEntity:       d.dataEntity,
		variablePropertyName: d.variablePropertyName,
	}
	return a, nil
}
