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

func createNEntityType(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.EntityType {
	items := make([]types.EntityType, n)
	for i := range items {
		items[i].Guid = strconv.Itoa(i)

		keeper.SetEntityType(ctx, items[i])
	}
	return items
}

func TestEntityTypeGet(t *testing.T) {
	keeper, ctx := keepertest.EavKeeper(t)
	items := createNEntityType(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetEntityType(ctx,
			item.Guid,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestEntityTypeRemove(t *testing.T) {
	keeper, ctx := keepertest.EavKeeper(t)
	items := createNEntityType(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEntityType(ctx,
			item.Guid,
		)
		_, found := keeper.GetEntityType(ctx,
			item.Guid,
		)
		require.False(t, found)
	}
}

func TestEntityTypeGetAll(t *testing.T) {
	keeper, ctx := keepertest.EavKeeper(t)
	items := createNEntityType(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllEntityType(ctx)),
	)
}
