package entity

type IDataPropertyModifier[T PropertyTypes] interface {
	NewFromThisData(
		executionVariables Entity,
		managerPropertyId *ManagerPropertyId,
	) (IPropertyModifier[T], error)
}
