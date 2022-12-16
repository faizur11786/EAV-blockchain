package eav

import (
	"belShare/x/eav/keeper"
	"belShare/x/eav/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the user
	for _, elem := range genState.UserList {
		k.SetUser(ctx, elem)
	}
	// Set all the merchant
	for _, elem := range genState.MerchantList {
		k.SetMerchant(ctx, elem)
	}
	// Set all the entityType
	for _, elem := range genState.EntityTypeList {
		k.SetEntityType(ctx, elem)
	}
	// Set all the attribute
	for _, elem := range genState.AttributeList {
		k.SetAttribute(ctx, elem)
	}
	// Set all the value
	for _, elem := range genState.ValueList {
		k.SetValue(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.UserList = k.GetAllUser(ctx)
	genesis.MerchantList = k.GetAllMerchant(ctx)
	genesis.EntityTypeList = k.GetAllEntityType(ctx)
	genesis.AttributeList = k.GetAllAttribute(ctx)
	genesis.ValueList = k.GetAllValue(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
