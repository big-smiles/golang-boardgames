package entity

type Predicate func(executionVariables Entity, entity Entity) (bool, error)

func filterEntities(executionVariables Entity, entities []*Entity, predicate Predicate) (*[]Entity, error) {
	count := 0
	for _, entity := range entities {
		keep, err := predicate(executionVariables, *entity)
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
		keep, err := predicate(executionVariables, *entity)
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
