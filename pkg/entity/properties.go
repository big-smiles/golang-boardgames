package entity

import (
	"fmt"
)

type PropertyTypes interface {
	int | string | bool | Id | []Id
}
type mapProperties[T PropertyTypes] map[PropertyId[T]]*property[T]
type properties struct {
	intProperties           *propertiesTyped[int]
	stringProperties        *propertiesTyped[string]
	boolProperties          *propertiesTyped[bool]
	entityIdProperties      *propertiesTyped[Id]
	arrayEntityIdProperties *propertiesTyped[[]Id]
}

func newProperties(m *ManagerPropertyId, d DataProperties) (*properties, error) {
	intProperties, err := newPropertiesTyped[int](m, d.IntProperties)
	if err != nil {
		return nil, err
	}
	stringProperties, err := newPropertiesTyped[string](m, d.StringProperties)
	if err != nil {
		return nil, err
	}
	boolProperties, err := newPropertiesTyped[bool](m, d.BoolProperties)
	if err != nil {
		return nil, err
	}
	entityIdProperties, err := newPropertiesTyped[Id](m, d.EntityIdProperties)
	if err != nil {
		return nil, err
	}
	arrayEntityIdProperties, err := newPropertiesTyped[[]Id](m, d.ArrayEntityIdProperties)
	if err != nil {
		return nil, err
	}
	return &properties{
		intProperties:           intProperties,
		stringProperties:        stringProperties,
		boolProperties:          boolProperties,
		entityIdProperties:      entityIdProperties,
		arrayEntityIdProperties: arrayEntityIdProperties,
	}, nil
}

/****************************MODIFIERS**************************************************/
func (p *properties) addModifier(m PropertiesModifier) error {
	err := p.intProperties.addModifier(m.intModifiers)
	if err != nil {
		return err
	}

	err = p.boolProperties.addModifier(m.boolModifiers)
	if err != nil {
		return err
	}

	err = p.stringProperties.addModifier(m.stringModifiers)
	if err != nil {
		return err
	}

	err = p.entityIdProperties.addModifier(m.entityIdModifiers)
	if err != nil {
		return err
	}
	err = p.arrayEntityIdProperties.addModifier(m.arrayEntityIdModifiers)
	if err != nil {
		return err
	}

	return nil
}

/****************************Output**************************************************/
//TODO: maybe move this to the outputProperties class
func (p *properties) getOutput() (*OutputProperties, error) {
	intProperties, err := p.intProperties.getOutput()
	if err != nil {
		return nil, err
	}

	stringProperties, err := p.stringProperties.getOutput()
	if err != nil {
		return nil, err
	}

	boolProperties, err := p.boolProperties.getOutput()
	if err != nil {
		return nil, err
	}

	entityIdProperties, err := p.entityIdProperties.getOutput()
	if err != nil {
		return nil, err
	}

	arrayIdProperties, err := p.arrayEntityIdProperties.getOutput()
	if err != nil {
		return nil, err
	}
	return NewOutputProperties(
		intProperties,
		stringProperties,
		boolProperties,
		entityIdProperties,
		arrayIdProperties,
	)
}

/********************************Properties**********************************************/
//getTypedProperties
func getTypedProperties[T PropertyTypes](p *properties) (*propertiesTyped[T], error) {
	var a T
	switch a2 := any(a).(type) {
	case int:
		return any(p.intProperties).(*propertiesTyped[T]), nil
	case string:
		return any(p.stringProperties).(*propertiesTyped[T]), nil
	case bool:
		return any(p.boolProperties).(*propertiesTyped[T]), nil
	case Id:
		return any(p.entityIdProperties).(*propertiesTyped[T]), nil
	case []Id:
		return any(p.arrayEntityIdProperties).(*propertiesTyped[T]), nil
	default:
		return nil, fmt.Errorf("unexpected type=%T", a2)
	}
}
