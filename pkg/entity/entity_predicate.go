package entity

type Predicate func(executionVariables Entity, managerPropertiesId *ManagerPropertyId, entity Entity) (bool, error)

func filterEntities(executionVariables Entity, managerPropertiesId *ManagerPropertyId, entities []*Entity, predicate Predicate) (*[]Entity, error) {
	count := 0
	for _, entity := range entities {
		keep, err := predicate(executionVariables, managerPropertiesId, *entity)
		if err != nil {
			return nil, err
		}
		if keep {
			count++
		}
	}
	result := make([]Entity, count)
	count = 0
	for _, entity := range entities {
		keep, err := predicate(executionVariables, managerPropertiesId, *entity)
		if err != nil {
			return nil, err
		}
		if keep {
			result[count] = *entity
			count++
		}
	}
	return &result, nil
}
