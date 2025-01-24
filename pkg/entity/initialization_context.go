package entity

type IInitializationContext interface {
	GetManagerEntityId() *ManagerEntityId
	GetManagerPropertyId() *ManagerPropertyId
}
