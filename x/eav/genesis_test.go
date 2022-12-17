package eav_test

import (
	"testing"

	keepertest "belShare/testutil/keeper"
	"belShare/testutil/nullify"
	"belShare/x/eav"
	"belShare/x/eav/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		UserList: []types.User{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		MerchantList: []types.Merchant{
			{
				Guid:    "0",
				Creator: "0",
			},
			{
				Guid:    "1",
				Creator: "1",
			},
		},
		EntityTypeList: []types.EntityType{
			{
				Guid: "0",
			},
			{
				Guid: "1",
			},
		},
		AttributeList: []types.Attribute{
			{
				Guid: "0",
			},
			{
				Guid: "1",
			},
		},
		ValueList: []types.Value{
			{
				EntityId:    "0",
				AttributeId: "0",
			},
			{
				EntityId:    "1",
				AttributeId: "1",
			},
		},
		MerchantNewList: []types.MerchantNew{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EavKeeper(t)
	eav.InitGenesis(ctx, *k, genesisState)
	got := eav.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.UserList, got.UserList)
	require.ElementsMatch(t, genesisState.MerchantList, got.MerchantList)
	require.ElementsMatch(t, genesisState.EntityTypeList, got.EntityTypeList)
	require.ElementsMatch(t, genesisState.AttributeList, got.AttributeList)
	require.ElementsMatch(t, genesisState.ValueList, got.ValueList)
	require.ElementsMatch(t, genesisState.MerchantNewList, got.MerchantNewList)
	// this line is used by starport scaffolding # genesis/test/assert
}
