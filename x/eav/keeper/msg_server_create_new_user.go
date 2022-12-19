package keeper

import (
	"context"

	"belShare/x/eav/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateNewUser(goCtx context.Context, msg *types.MsgCreateNewUser) (*types.MsgCreateNewUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateNewUserResponse{}, nil
}
