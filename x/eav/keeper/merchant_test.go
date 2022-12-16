package keeper_test

import (
	"strconv"
	"testing"

	keepertest "belShare/testutil/keeper"
	"belShare/testutil/nullify"
	"belShare/x/eav/keeper"
	"belShare/x/eav/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNMerchant(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Merchant {
	items := make([]types.Merchant, n)
	for i := range items {
		items[i].Guid = strconv.Itoa(i)
		items[i].Creator = strconv.Itoa(i)

		keeper.SetMerchant(ctx, items[i])
	}
	return items
}

func TestMerchantGet(t *testing.T) {
	keeper, ctx := keepertest.EavKeeper(t)
	items := createNMerchant(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetMerchant(ctx,
			item.Guid,
			item.Creator,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestMerchantRemove(t *testing.T) {
	keeper, ctx := keepertest.EavKeeper(t)
	items := createNMerchant(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMerchant(ctx,
			item.Guid,
			item.Creator,
		)
		_, found := keeper.GetMerchant(ctx,
			item.Guid,
			item.Creator,
		)
		require.False(t, found)
	}
}

func TestMerchantGetAll(t *testing.T) {
	keeper, ctx := keepertest.EavKeeper(t)
	items := createNMerchant(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMerchant(ctx)),
	)
}
