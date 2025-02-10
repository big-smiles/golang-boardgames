package entity

import (
	"errors"
	"fmt"
)

type ManagerEntity struct {
	entitiesById      map[Id]*Entity
	allEntities       []*Entity
	archetypesByName  map[string]DataProperties
	managerEntityId   *ManagerEntityId
	managerPropertyId *ManagerPropertyId
}

func NewManagerEntity() (*ManagerEntity, error) {
	return &ManagerEntity{
		entitiesById:     make(map[Id]*Entity),
		allEntities:      make([]*Entity, 0),
		archetypesByName: make(map[string]DataProperties),
	}, nil
}

func (m *ManagerEntity) Initialize(ctx IInitializationContext) error {
	m.managerEntityId = ctx.GetManagerEntityId()
	m.managerPropertyId = ctx.GetManagerPropertyId()
	return nil
}

func (m *ManagerEntity) New(executionVariables Entity, managerPropertyId *ManagerPropertyId, d DataEntity) (*Entity, error) {
	entity, err := newEntity(executionVariables, m.managerEntityId, managerPropertyId, d)
	if err != nil {
		return nil, err
	}
	_, ok := m.entitiesById[entity.Id]
	if ok {
		return nil, errors.New(fmt.Sprintf("Entity with id=%d already exists", entity.Id))
	}
	m.entitiesById[entity.Id] = entity
	m.allEntities = append(m.allEntities, entity)
	return entity, nil
}
func (m *ManagerEntity) NewExecutionVariable(d DataEntity) (*Entity, error) {
	if m.managerEntityId == nil {
		return nil, errors.New(fmt.Sprintf("Manager Entity Id not initialized"))
	}
	entity, err := newEntityForExecutionVariable(m.managerPropertyId, d, m.managerEntityId)
	if err != nil {
		return nil, err
	}
	_, ok := m.entitiesById[entity.Id]
	if ok {
		return nil, errors.New(fmt.Sprintf("Entity with id=%d already exists", entity.Id))
	}
	m.entitiesById[entity.Id] = entity
	m.allEntities = append(m.allEntities, entity)
	return entity, nil
}
func (m *ManagerEntity) FindById(id Id) (*Entity, error) {
	e, ok := m.entitiesById[id]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Entity with id=%d not found", id))
	}
	return e, nil
}

func (m *ManagerEntity) GetOutput(output *[]OutputEntity) error {
	i := 0
	for _, e := range m.allEntities {
		err := e.getOutput(output, i)
		if err != nil {
			return err
		}
		if e.isExecutionVariable == false {
			i++
		}
	}
	return nil
}
func (m *ManagerEntity) GetOutputAmount() int {
	count := 0
	for _, e := range m.allEntities {
		if e.isExecutionVariable {
			count++
		}
	}
	return len(m.entitiesById) - count

}

func (m *ManagerEntity) GetFiltered(executionVariables Entity, predicate Predicate) ([]Id, error) {
	entities, err := filterEntities(executionVariables, m.managerPropertyId, m.allEntities, predicate)
	if err != nil {
		return nil, err
	}
	result := make([]Id, len(*entities))
	for i, entity := range *entities {
		result[i] = entity.Id
	}
	return result, nil
}
