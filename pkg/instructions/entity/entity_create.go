package instruction_entity

import (
	"fmt"
	"github.com/big-smiles/golang-boardgames/pkg/entity"
	"github.com/big-smiles/golang-boardgames/pkg/instruction"
)

type instructionCreateEntity struct {
	nameDataEntity entity.NameDataEntity
}

func (i instructionCreateEntity) Execute(ctx instruction.ExecutionContext) error {
	data, err := ctx.Performer.Entity.GetData(i.nameDataEntity)
	if err != nil {
		return err
	}
	id, err := ctx.Performer.Entity.Create(ctx.ExecutionVariables, data)
	fmt.Printf("created entity id=%d", id)
	if err != nil {
		return err
	}
	return nil
}

func newInstructionCreateEntity(d DataInstructionCreateEntity) (*instructionCreateEntity, error) {
	a := &instructionCreateEntity{
		nameDataEntity: d.dataEntity,
	}
	return a, nil
}
