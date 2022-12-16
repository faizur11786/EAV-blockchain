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

func createNAttribute(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Attribute {
	items := make([]types.Attribute, n)
	for i := range items {
		items[i].Guid = strconv.Itoa(i)

		keeper.SetAttribute(ctx, items[i])
	}
	return items
}

func TestAttributeGet(t *testing.T) {
	keeper, ctx := keepertest.EavKeeper(t)
	items := createNAttribute(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAttribute(ctx,
			item.Guid,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAttributeRemove(t *testing.T) {
	keeper, ctx := keepertest.EavKeeper(t)
	items := createNAttribute(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAttribute(ctx,
			item.Guid,
		)
		_, found := keeper.GetAttribute(ctx,
			item.Guid,
		)
		require.False(t, found)
	}
}

func TestAttributeGetAll(t *testing.T) {
	keeper, ctx := keepertest.EavKeeper(t)
	items := createNAttribute(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAttribute(ctx)),
	)
}
