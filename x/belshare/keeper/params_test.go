package keeper_test

import (
	"testing"

	testkeeper "belShare/testutil/keeper"
	"belShare/x/belshare/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BelshareKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
