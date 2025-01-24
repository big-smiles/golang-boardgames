package instruction

type Performer struct {
	Entity        *Entity
	Output        *Output
	Interaction   *Interaction
	ValueResolver *ValueResolver
}

func NewPerformer() (*Performer, error) {
	entityPerformer, err := NewPerformerEntity()
	if err != nil {
		return nil, err
	}

	outputPerformer, err := NewPerformerOutput()
	if err != nil {
		return nil, err
	}

	interaction, err := NewPerformerInteraction()
	if err != nil {
		return nil, err
	}

	valueResolver, err := NewPerformerValueResolver()
	if err != nil {
		return nil, err
	}
	return &Performer{
		Entity:        entityPerformer,
		Output:        outputPerformer,
		Interaction:   interaction,
		ValueResolver: valueResolver,
	}, nil
}

func (p *Performer) Initialize(ctx InitializationContext) error {
	err := p.Entity.Initialize(ctx)
	if err != nil {
		return err
	}

	err = p.Output.Initialize(ctx)
	if err != nil {
		return err
	}

	err = p.Interaction.Initialize(ctx)
	if err != nil {
		return err
	}

	err = p.ValueResolver.Initialize(ctx)
	if err != nil {
		return err
	}
	return nil
}
