package instruction

import (
	"errors"
	"github.com/big-smiles/golang-boardgames/pkg/entity"
	resolveValueConstant "github.com/big-smiles/golang-boardgames/pkg/resolve_value/constant"
	ValueModifierCommon "github.com/big-smiles/golang-boardgames/pkg/value_modifier/common"
)

type Entity struct {
	managerEntity     *entity.ManagerEntity
	managerEntityId   *entity.ManagerEntityId
	managerPropertyId *entity.ManagerPropertyId
	managerEntityData *entity.ManagerData
}

func NewPerformerEntity() (*Entity, error) {
	return &Entity{}, nil
}

func (p *Entity) GetId(executionVariables entity.Entity, d entity.DataId) (entity.Id, error) {
	if p.managerEntityId == nil {
		return 0, errors.New("in Entity managerEntityId is nil")
	}
	id, err := p.managerEntityId.GetId(executionVariables, p.managerPropertyId, d)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *Entity) FilterEntitiesIntoVariable(
	executionVariables entity.Entity,
	predicate entity.Predicate,
	namePropertyId entity.NamePropertyId[[]entity.Id],
) error {
	filteredIds, err := p.managerEntity.GetFiltered(executionVariables, predicate)
	if err != nil {
		return err
	}
	valueResolver := resolveValueConstant.NewResolveConstant[[]entity.Id](filteredIds)
	propertyDataModifier, err :=
		ValueModifierCommon.NewDataModifierSetValue[[]entity.Id](valueResolver)
	if err != nil {
		return err
	}
	mapDataModifierProperties := make(entity.MapDataModifierProperties[[]entity.Id], 1)

	mapDataModifierProperties[namePropertyId] = propertyDataModifier

	dataPropertiesModifier := entity.DataPropertiesModifier{
		ArrayEntityIdModifiers: mapDataModifierProperties,
	}

	dataEntityModifier, err := entity.NewDataEntityModifier(dataPropertiesModifier)
	if err != nil {
		return err
	}
	err = p.AddModifier(executionVariables, []entity.Id{executionVariables.Id}, *dataEntityModifier)
	if err != nil {
		return err
	}

	return nil
}

func (p *Entity) GetData(name entity.NameDataEntity) (entity.DataEntity, error) {
	if p.managerEntityData == nil {
		return entity.DataEntity{}, errors.New("in Entity managerEntityData is nil")
	}
	data, err := p.managerEntityData.Get(name)
	if err != nil {
		return entity.DataEntity{}, err
	}

	return data, nil
}

func (p *Entity) Create(executionVariables entity.Entity, d entity.DataEntity) (entity.Id, error) {
	if p.managerEntity == nil {
		return 0, errors.New("manager entity is nil")
	}
	ent, err := p.managerEntity.New(executionVariables, p.managerPropertyId, d)
	if err != nil {
		return 0, err
	}
	return ent.Id, nil
}

func (p *Entity) Get(executionVariables entity.Entity, d entity.DataId) (*entity.Entity, error) {
	if p.managerEntityId == nil {
		return nil, errors.New("in Entity managerEntityId is nil")
	}
	id, err := p.managerEntityId.GetId(executionVariables, p.managerPropertyId, d)
	if err != nil {
		return nil, err
	}
	ent, err := p.managerEntity.FindById(id)
	if err != nil {
		return nil, err
	}
	return ent, nil
}

func (p *Entity) GetById(id entity.Id) (*entity.Entity, error) {
	ent, err := p.managerEntity.FindById(id)
	if err != nil {
		return nil, err
	}
	return ent, nil
}

func (p *Entity) AddModifier(executionVariables entity.Entity, target []entity.Id, d entity.DataModifier) error {
	if p.managerEntity == nil {
		return errors.New("manager entity is nil")
	}
	em, err := entity.NewEntityModifier(executionVariables, p.managerPropertyId, d)
	if err != nil {
		return err
	}
	for _, id := range target {
		e, err := p.managerEntity.FindById(id)
		if err != nil {
			return err
		}
		err = e.AddModifier(*em)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetValueFromEntity[T entity.PropertyTypes](
	performerEntity Entity,
	e entity.Entity,
	name entity.NamePropertyId[T],
) (T, error) {
	var zero T

	typedManager, err := entity.GetManagerTypedPropertyId[T](performerEntity.managerPropertyId)
	if err != nil {
		return zero, err
	}

	id, err := typedManager.GetId(name)
	if err != nil {
		return zero, err
	}

	value, err := entity.GetValueFromEntity(e, id)
	if err != nil {
		return zero, err
	}

	return value, nil
}

func (p *Entity) Initialize(ctx InitializationContext) error {
	p.managerEntity = ctx.GetManagerEntity()
	p.managerEntityId = ctx.GetManagerEntityId()
	p.managerPropertyId = ctx.GetManagerPropertyId()
	p.managerEntityData = ctx.GetManagerEntityData()
	return nil
}
