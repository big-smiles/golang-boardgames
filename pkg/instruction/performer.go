package instruction

type Performer struct {
	Entity        *Entity
	Output        *Output
	Interaction   *Interaction
	ValueResolver *ValueResolver
	Phase         *Phase
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

	phase, err := NewPerformerPhase()
	if err != nil {
		return nil, err
	}
	return &Performer{
		Entity:        entityPerformer,
		Output:        outputPerformer,
		Interaction:   interaction,
		ValueResolver: valueResolver,
		Phase:         phase,
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

	err = p.Phase.Initialize(ctx)
	if err != nil {
		return err
	}

	return nil
}
