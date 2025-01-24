package instruction

import (
	"errors"
	"github.com/big-smiles/golang-boardgames/pkg/entity"
	resolveValueConstant "github.com/big-smiles/golang-boardgames/pkg/resolve_value/constant"
	ValueModifierCommon "github.com/big-smiles/golang-boardgames/pkg/value_modifier/common"
)

type ManagerInstruction struct {
	stack         *instructionStack
	executing     bool
	performer     *Performer
	managerEntity *entity.ManagerEntity
}

func NewManagerInstruction() (*ManagerInstruction, error) {
	s, err := newInstructionStack()
	if err != nil {
		return nil, err
	}

	return &ManagerInstruction{
		stack:     s,
		executing: false,
		performer: nil,
	}, nil
}

func (m *ManagerInstruction) AddInstruction(
	i Instruction,
	selectedEntities []entity.Id,
) error {
	w, err := newWrapperInstruction(i, selectedEntities)
	if err != nil {
		return err
	}

	m.stack.push(*w)

	err = m.executeLoop()
	if err != nil {
		return err
	}

	return nil
}
func (m *ManagerInstruction) executeLoop() error {
	if m.executing {
		return errors.New("already executing")
	}
	m.executing = true
	for w, ok := m.stack.pop(); ok == true; w, ok = m.stack.pop() {
		ctx, err := m.buildExecutionContext(w.selectedEntities)
		if err != nil {
			return err
		}
		err = w.instruction.Execute(*ctx)
		if err != nil {
			return err
		}
	}
	m.executing = false

	return nil
}
func (m *ManagerInstruction) buildExecutionContext(selectedEntities []entity.Id) (*ExecutionContext, error) {
	if m.managerEntity == nil {
		return nil, errors.New("manager instruction does not have manager entity")
	}
	executionVariableData, err := getExecutionVariablesData()
	if err != nil {
		return nil, err
	}
	executionVariable, err := m.managerEntity.NewExecutionVariable(*executionVariableData)
	if err != nil {
		return nil, err
	}
	ctx, err := newExecutionContext(m.performer, *executionVariable)
	resolveSelectedEntities, err := resolveValueConstant.NewResolveConstant[[]entity.Id](selectedEntities)
	if err != nil {
		return nil, err
	}
	dataModifierSetValue, err := ValueModifierCommon.NewDataModifierSetValue[[]entity.Id](resolveSelectedEntities)
	if err != nil {
		return nil, err
	}
	mapDataModifierProperties := make(entity.MapDataModifierProperties[[]entity.Id], 1)
	mapDataModifierProperties[SELECTED_ENTITIES] = dataModifierSetValue
	dataPropertiesModifier, err := entity.NewDataPropertiesModifier(
		nil,
		nil,
		nil,
		nil,
		&mapDataModifierProperties,
	)
	if err != nil {
		return nil, err
	}
	dataEntityModifier, err := entity.NewDataEntityModifier(*dataPropertiesModifier)
	if err != nil {
		return nil, err
	}

	err = m.performer.Entity.AddModifier(
		ctx.ExecutionVariables,
		[]entity.Id{ctx.ExecutionVariables.Id},
		*dataEntityModifier,
	)
	if err != nil {
		return nil, err
	}
	return ctx, nil
}

func (m *ManagerInstruction) Initialize(ctx InitializationContext) error {
	m.performer = ctx.GetPerformer()
	m.managerEntity = ctx.GetManagerEntity()
	return nil
}
