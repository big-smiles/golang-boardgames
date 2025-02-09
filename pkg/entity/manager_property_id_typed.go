package entity

type ManagerTypedPropertyId[T PropertyTypes] struct {
	counter  PropertyId[T]
	idByName mapIdProperty[T]
}

func NewManagerTypedPropertyId[T PropertyTypes]() *ManagerTypedPropertyId[T] {
	return &ManagerTypedPropertyId[T]{
		counter:  0,
		idByName: make(mapIdProperty[T]),
	}
}
func (m *ManagerTypedPropertyId[T]) Initialize(_ IInitializationContext) error {

	return nil
}

func (m *ManagerTypedPropertyId[T]) GetId(name NamePropertyId[T]) (PropertyId[T], error) {
	if name == "" {
		m.counter++
		newId := m.counter
		return newId, nil
	}
	val, ok := m.idByName[name]
	if !ok {
		m.counter++
		newId := m.counter
		m.idByName[name] = newId
		return newId, nil
	}
	return val, nil
}
