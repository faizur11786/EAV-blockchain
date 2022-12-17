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

func createNMerchantNew(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.MerchantNew {
	items := make([]types.MerchantNew, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetMerchantNew(ctx, items[i])
	}
	return items
}

func TestMerchantNewGet(t *testing.T) {
	keeper, ctx := keepertest.EavKeeper(t)
	items := createNMerchantNew(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetMerchantNew(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestMerchantNewRemove(t *testing.T) {
	keeper, ctx := keepertest.EavKeeper(t)
	items := createNMerchantNew(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMerchantNew(ctx,
			item.Address,
		)
		_, found := keeper.GetMerchantNew(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestMerchantNewGetAll(t *testing.T) {
	keeper, ctx := keepertest.EavKeeper(t)
	items := createNMerchantNew(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMerchantNew(ctx)),
	)
}
