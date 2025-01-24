package entity

type IValueResolver[T any] interface {
	Resolve(
		executionVariables Entity,
		managerPropertyId *ManagerPropertyId,
	) (T, error)
}
