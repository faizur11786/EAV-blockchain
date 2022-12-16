package belshare_test

import (
	"testing"

	keepertest "belShare/testutil/keeper"
	"belShare/testutil/nullify"
	"belShare/x/belshare"
	"belShare/x/belshare/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BelshareKeeper(t)
	belshare.InitGenesis(ctx, *k, genesisState)
	got := belshare.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
