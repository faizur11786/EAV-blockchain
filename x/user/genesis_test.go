package user_test

import (
	"testing"

	keepertest "belShare/testutil/keeper"
	"belShare/testutil/nullify"
	"belShare/x/user"
	"belShare/x/user/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		UserList: []types.User{
			{
				Creator: "0",
				Guid:    "0",
			},
			{
				Creator: "1",
				Guid:    "1",
			},
		},
		XxxList: []types.Xxx{
			{
				Address: "0",
				Creator: "0",
			},
			{
				Address: "1",
				Creator: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.UserKeeper(t)
	user.InitGenesis(ctx, *k, genesisState)
	got := user.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.UserList, got.UserList)
	require.ElementsMatch(t, genesisState.XxxList, got.XxxList)
	// this line is used by starport scaffolding # genesis/test/assert
}
