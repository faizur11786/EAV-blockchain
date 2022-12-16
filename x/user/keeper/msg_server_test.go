package keeper_test

import (
	"context"
	"testing"

	keepertest "belShare/testutil/keeper"
	"belShare/x/user/keeper"
	"belShare/x/user/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.UserKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
