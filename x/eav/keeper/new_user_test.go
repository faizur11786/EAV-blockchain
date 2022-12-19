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

func createNNewUser(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.NewUser {
	items := make([]types.NewUser, n)
	for i := range items {
		items[i].Guid = strconv.Itoa(i)

		keeper.SetNewUser(ctx, items[i])
	}
	return items
}

func TestNewUserGet(t *testing.T) {
	keeper, ctx := keepertest.EavKeeper(t)
	items := createNNewUser(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNewUser(ctx,
			item.Guid,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestNewUserRemove(t *testing.T) {
	keeper, ctx := keepertest.EavKeeper(t)
	items := createNNewUser(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNewUser(ctx,
			item.Guid,
		)
		_, found := keeper.GetNewUser(ctx,
			item.Guid,
		)
		require.False(t, found)
	}
}

func TestNewUserGetAll(t *testing.T) {
	keeper, ctx := keepertest.EavKeeper(t)
	items := createNNewUser(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNewUser(ctx)),
	)
}
