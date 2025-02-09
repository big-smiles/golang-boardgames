package ValueModifierCommon

type ModifierSetValue[T any] struct {
	value T
}

func NewModifierSetValue[T any](value T) (*ModifierSetValue[T], error) {
	return &ModifierSetValue[T]{
		value: value,
	}, nil
}

func (m *ModifierSetValue[T]) Modify(_ T) (newValue T, err error) {
	return m.value, nil
}
