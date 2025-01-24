package entity

type IPropertyModifier[T PropertyTypes] interface {
	Modify(prevValue T) (newValue T, err error)
}
