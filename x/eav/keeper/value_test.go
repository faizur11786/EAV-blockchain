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

func createNValue(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Value {
	items := make([]types.Value, n)
	for i := range items {
		items[i].EntityId = strconv.Itoa(i)
		items[i].AttributeId = strconv.Itoa(i)

		keeper.SetValue(ctx, items[i])
	}
	return items
}

func TestValueGet(t *testing.T) {
	keeper, ctx := keepertest.EavKeeper(t)
	items := createNValue(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetValue(ctx,
			item.EntityId,
			item.AttributeId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestValueRemove(t *testing.T) {
	keeper, ctx := keepertest.EavKeeper(t)
	items := createNValue(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveValue(ctx,
			item.EntityId,
			item.AttributeId,
		)
		_, found := keeper.GetValue(ctx,
			item.EntityId,
			item.AttributeId,
		)
		require.False(t, found)
	}
}

func TestValueGetAll(t *testing.T) {
	keeper, ctx := keepertest.EavKeeper(t)
	items := createNValue(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllValue(ctx)),
	)
}
