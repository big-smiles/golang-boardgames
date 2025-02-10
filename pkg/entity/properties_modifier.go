package entity

type mapModifierProperties[T PropertyTypes] map[PropertyId[T]]IPropertyModifier[T]
type PropertiesModifier struct {
	intModifiers           mapModifierProperties[int]
	stringModifiers        mapModifierProperties[string]
	boolModifiers          mapModifierProperties[bool]
	entityIdModifiers      mapModifierProperties[Id]
	arrayEntityIdModifiers mapModifierProperties[[]Id]
}

func NewPropertiesModifier(
	executionVariables Entity,
	m *ManagerPropertyId,
	d DataPropertiesModifier,
) (*PropertiesModifier, error) {
	intModifiers, err := _newPropertiesMapData[int](executionVariables, m, d.IntModifiers)
	if err != nil {
		return nil, err
	}

	stringModifiers, err := _newPropertiesMapData[string](executionVariables, m, d.StringModifiers)
	if err != nil {
		return nil, err
	}

	boolModifiers, err := _newPropertiesMapData[bool](executionVariables, m, d.BoolModifiers)
	if err != nil {
		return nil, err
	}

	entityIdModifiers, err := _newPropertiesMapData[Id](executionVariables, m, d.EntityIdModifiers)
	if err != nil {
		return nil, err
	}
	arrayEntityIdModifiers, err := _newPropertiesMapData[[]Id](executionVariables, m, d.ArrayEntityIdModifiers)
	if err != nil {
		return nil, err
	}
	return &PropertiesModifier{
		intModifiers:           intModifiers,
		stringModifiers:        stringModifiers,
		boolModifiers:          boolModifiers,
		entityIdModifiers:      entityIdModifiers,
		arrayEntityIdModifiers: arrayEntityIdModifiers,
	}, nil
}

func _newPropertiesMapData[T PropertyTypes](
	executionVariables Entity,
	m *ManagerPropertyId,
	d MapDataModifierProperties[T],
) (mapModifierProperties[T], error) {
	r := make(mapModifierProperties[T], len(d))
	for k, v := range d {
		pm, err := v.NewFromThisData(executionVariables, m)
		if err != nil {
			return nil, err
		}
		managerTypedPropertyId, err := GetManagerTypedPropertyId[T](m)
		if err != nil {
			return nil, err
		}
		id, err := managerTypedPropertyId.GetId(k)
		if err != nil {
			return nil, err
		}

		r[id] = pm
	}
	return r, nil
}
