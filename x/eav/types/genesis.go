package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		UserList:        []User{},
		MerchantList:    []Merchant{},
		EntityTypeList:  []EntityType{},
		AttributeList:   []Attribute{},
		ValueList:       []Value{},
		MerchantNewList: []MerchantNew{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in user
	userIndexMap := make(map[string]struct{})

	for _, elem := range gs.UserList {
		index := string(UserKey(elem.Index))
		if _, ok := userIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for user")
		}
		userIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in merchant
	merchantIndexMap := make(map[string]struct{})

	for _, elem := range gs.MerchantList {
		index := string(MerchantKey(elem.Guid, elem.Creator))
		if _, ok := merchantIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for merchant")
		}
		merchantIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in entityType
	entityTypeIndexMap := make(map[string]struct{})

	for _, elem := range gs.EntityTypeList {
		index := string(EntityTypeKey(elem.Guid))
		if _, ok := entityTypeIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for entityType")
		}
		entityTypeIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in attribute
	attributeIndexMap := make(map[string]struct{})

	for _, elem := range gs.AttributeList {
		index := string(AttributeKey(elem.Guid))
		if _, ok := attributeIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for attribute")
		}
		attributeIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in value
	valueIndexMap := make(map[string]struct{})

	for _, elem := range gs.ValueList {
		index := string(ValueKey(elem.EntityId, elem.AttributeId))
		if _, ok := valueIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for value")
		}
		valueIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in merchantNew
	merchantNewIndexMap := make(map[string]struct{})

	for _, elem := range gs.MerchantNewList {
		index := string(MerchantNewKey(elem.Address))
		if _, ok := merchantNewIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for merchantNew")
		}
		merchantNewIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
