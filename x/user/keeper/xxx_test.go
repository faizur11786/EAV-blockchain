package keeper_test

import (
	"strconv"
	"testing"

	keepertest "belShare/testutil/keeper"
	"belShare/testutil/nullify"
	"belShare/x/user/keeper"
	"belShare/x/user/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNXxx(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Xxx {
	items := make([]types.Xxx, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)
		items[i].Creator = strconv.Itoa(i)

		keeper.SetXxx(ctx, items[i])
	}
	return items
}

func TestXxxGet(t *testing.T) {
	keeper, ctx := keepertest.UserKeeper(t)
	items := createNXxx(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetXxx(ctx,
			item.Address,
			item.Creator,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestXxxRemove(t *testing.T) {
	keeper, ctx := keepertest.UserKeeper(t)
	items := createNXxx(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveXxx(ctx,
			item.Address,
			item.Creator,
		)
		_, found := keeper.GetXxx(ctx,
			item.Address,
			item.Creator,
		)
		require.False(t, found)
	}
}

func TestXxxGetAll(t *testing.T) {
	keeper, ctx := keepertest.UserKeeper(t)
	items := createNXxx(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllXxx(ctx)),
	)
}
