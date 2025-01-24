package entity

import (
	"fmt"
)

type propertiesTyped[T PropertyTypes] struct {
	properties mapProperties[T]
}

func newPropertiesTyped[T PropertyTypes](m *ManagerPropertyId, d DataPropertiesTyped[T]) (*propertiesTyped[T], error) {
	props := make(mapProperties[T], len(d))
	for _, name := range d {
		managerTypedPropertyId, err := GetManagerTypedPropertyId[T](m)
		if err != nil {
			return nil, err
		}
		id, err := managerTypedPropertyId.GetId(name)
		prop, err := newProperty[T](id)
		if err != nil {
			return nil, err
		}
		props[id] = prop
	}
	return &propertiesTyped[T]{
		properties: props,
	}, nil
}
func (p *propertiesTyped[T]) addProperty(id PropertyId[T]) error {
	prop, err := newProperty[T](id)
	if err != nil {
		return err
	}
	p.properties[id] = prop
	return nil
}
func (p *propertiesTyped[T]) addModifier(m mapModifierProperties[T]) error {
	for key, modifier := range m {
		prop, ok := p.properties[key]
		if !ok {
			err := p.addProperty(key)
			if err != nil {
				return err
			}
			prop, ok = p.properties[key]
			if !ok {
				return fmt.Errorf("tried to add property for modifier but it wasnt added key=%d", key)
			}
		}
		err := prop.AddModifier(modifier)
		if err != nil {
			return err
		}
	}
	return nil
}
func (p *propertiesTyped[T]) getOutput() (*mapOutputProperties[T], error) {
	ret := make(mapOutputProperties[T], len(p.properties))
	for key, prop := range p.properties {
		v := prop.Get()
		ret[key] = v
	}
	return &ret, nil
}
func (p *propertiesTyped[T]) GetProperty(id PropertyId[T]) (*property[T], error) {
	prop, ok := p.properties[id]
	if !ok {
		return prop, fmt.Errorf("property not found id=%d", id)
	}
	return prop, nil
}
