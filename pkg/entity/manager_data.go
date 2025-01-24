package entity

import (
	"fmt"
)

// NameDataEntity alias for string
//
// used by ManagerData to sort and provide DataEntity
type NameDataEntity string

// LibraryDataEntity alias for map[NameDataEntity]DataEntity
//
// used to initialize ManagerData
type LibraryDataEntity map[NameDataEntity]DataEntity

// ManagerData keeps all data to instantiate entities, sorted by a NameDataEntity
//
// # Usage
//
// - Initialize before using
//
// - Get receives a NameDataEntity and returns a DataEntity
type ManagerData struct {
	data map[NameDataEntity]DataEntity
}

// NewManagerData create a new instance of ManagerData, receives the initial data to be loaded
func NewManagerData(d map[NameDataEntity]DataEntity) (*ManagerData, error) {
	return &ManagerData{
		data: d,
	}, nil
}

// Initialize with the engine context, must be called before using
func (md *ManagerData) Initialize(_ IInitializationContext) error {
	return nil
}

// Get receives a NameDataEntity and returns a DataEntity
//
// Returns error if no DataEntity is registered for the give NameDataEntity
func (md *ManagerData) Get(n NameDataEntity) (DataEntity, error) {
	data, ok := md.data[n]
	if !ok {
		return DataEntity{}, fmt.Errorf("dataEntity not found name=%s", n)
	}
	return data, nil
}
