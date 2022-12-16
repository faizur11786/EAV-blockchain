package keeper

import (
	"context"

	"belShare/x/eav/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateUser(goCtx context.Context, msg *types.MsgCreateUser) (*types.MsgCreateUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create variable of type USER
	var user = types.User{
		Index:     msg.Creator,
		Creator:   msg.Address,
		CreatedAt: ctx.BlockHeight(),
	}

	// Write user information to the store
	k.SetUser(ctx, user)

	return &types.MsgCreateUserResponse{}, nil
}
