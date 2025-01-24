package entity

type property[T PropertyTypes] struct {
	id           PropertyId[T]
	initialValue T
	cachedValue  T
	modifiers    []IPropertyModifier[T]
}

// NewProperty returns the property common as the zero value
func newProperty[T PropertyTypes](id PropertyId[T]) (*property[T], error) {
	return &property[T]{
		id:        id,
		modifiers: make([]IPropertyModifier[T], 0),
	}, nil
}

func (p *property[T]) Get() T {
	return p.cachedValue
}

func (p *property[T]) AddModifier(m IPropertyModifier[T]) error {
	p.modifiers = append(p.modifiers, m)
	return p.recalculateValue()
}
func (p *property[T]) recalculateValue() error {
	p.cachedValue = p.initialValue
	for _, m := range p.modifiers {
		v, err := m.Modify(p.cachedValue)
		if err != nil {
			return err
		}
		p.cachedValue = v
	}
	return nil
}
