package entity

import (
	"errors"
)

type Entity struct {
	Id                  Id
	Name                NameEntityId
	properties          *properties
	isExecutionVariable bool
}

func newEntity(
	executionVariables Entity,
	managerEntityId *ManagerEntityId,
	managerPropertyId *ManagerPropertyId,
	d DataEntity,
) (*Entity, error) {
	if managerEntityId == nil {
		return nil, errors.New("managerEntityId is nil")
	}
	name, err := d.id.ResolverName.Resolve(executionVariables, managerPropertyId)
	if err != nil {
		return nil, err
	}
	id, err := managerEntityId.GetId(
		executionVariables,
		managerPropertyId,
		d.id,
	)
	if err != nil {
		return nil, err
	}

	p, err := newProperties(managerPropertyId, d.dataProperties)
	if err != nil {
		return nil, err
	}

	return &Entity{
		Id:                  id,
		Name:                name,
		properties:          p,
		isExecutionVariable: false,
	}, nil
}
func newEntityForExecutionVariable(
	managerPropertyId *ManagerPropertyId,
	d DataEntity,
	managerEntityId *ManagerEntityId,
) (*Entity, error) {
	if managerEntityId == nil {
		return nil, errors.New("managerEntityId is nil")
	}
	id, err := managerEntityId.getNextId()
	if err != nil {
		return nil, err
	}
	p, err := newProperties(managerPropertyId, d.dataProperties)
	return &Entity{
		Id:                  id,
		properties:          p,
		isExecutionVariable: true,
	}, nil
}
func (e Entity) getOutput(output *[]OutputEntity, i int) error {
	if e.isExecutionVariable {
		return nil
	}
	o, err := NewOutputEntity(e)
	if err != nil {
		return err
	}
	(*output)[i] = *o
	return nil

}
func (e Entity) AddModifier(m Modifier) error {
	err := e.properties.addModifier(m.propertiesModifier)
	if err != nil {
		return err
	}
	return nil
}

// GetValueFromEntity
// Errors: ErrorPropertyNotFound, GenericError when the typedProperty is not set
func GetValueFromEntity[T PropertyTypes](
	e Entity,
	id PropertyId[T],
) (T, error) {
	var zero T

	typedProperties, err := getTypedProperties[T](e.properties)
	if err != nil {
		return zero, err
	}

	p, err := typedProperties.GetProperty(id)
	if err != nil {
		return zero, err
	}

	val := p.Get()

	return val, nil
}
