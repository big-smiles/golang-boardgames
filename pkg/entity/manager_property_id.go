package entity

import "fmt"

type NamePropertyId[T PropertyTypes] string
type PropertyId[T PropertyTypes] uint64
type mapIdProperty[T PropertyTypes] map[NamePropertyId[T]]PropertyId[T]

type ManagerPropertyId struct {
	string        *ManagerTypedPropertyId[string]
	int           *ManagerTypedPropertyId[int]
	bool          *ManagerTypedPropertyId[bool]
	entityId      *ManagerTypedPropertyId[Id]
	arrayEntityId *ManagerTypedPropertyId[[]Id]
}

func NewManagerPropertyId() (*ManagerPropertyId, error) {
	return &ManagerPropertyId{
		string:        NewManagerTypedPropertyId[string](),
		int:           NewManagerTypedPropertyId[int](),
		bool:          NewManagerTypedPropertyId[bool](),
		entityId:      NewManagerTypedPropertyId[Id](),
		arrayEntityId: NewManagerTypedPropertyId[[]Id](),
	}, nil
}

func (m *ManagerPropertyId) Initialize(ctx IInitializationContext) error {
	err := m.string.Initialize(ctx)
	if err != nil {
		return err
	}
	err = m.int.Initialize(ctx)
	if err != nil {
		return err
	}
	err = m.bool.Initialize(ctx)
	if err != nil {
		return err
	}
	err = m.entityId.Initialize(ctx)
	if err != nil {
		return err
	}
	err = m.arrayEntityId.Initialize(ctx)
	if err != nil {
		return err
	}
	return nil
}

func GetManagerTypedPropertyId[T PropertyTypes](m *ManagerPropertyId) (*ManagerTypedPropertyId[T], error) {
	var a1 T
	switch a2 := any(a1).(type) {
	case int:
		return any(m.int).(*ManagerTypedPropertyId[T]), nil
	case string:
		return any(m.string).(*ManagerTypedPropertyId[T]), nil
	case bool:
		return any(m.bool).(*ManagerTypedPropertyId[T]), nil
	case Id:
		return any(m.entityId).(*ManagerTypedPropertyId[T]), nil
	case []Id:
		return any(m.arrayEntityId).(*ManagerTypedPropertyId[T]), nil
	default:
		return nil, fmt.Errorf("GetManagerTypedPropertyId[%s]: %s. type unimplemented", a2, a2)
	}
}

func (m *ManagerPropertyId) GetOutput() (OutputManagerPropertyId, error) {
	o := NewOutputManagerPropertyId(
		m.string.idByName,
		m.int.idByName,
		m.bool.idByName,
		m.entityId.idByName,
		m.arrayEntityId.idByName,
	)
	return *o, nil
}
