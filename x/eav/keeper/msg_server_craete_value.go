package keeper

import (
	"context"

	"belShare/x/eav/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CraeteValue(goCtx context.Context, msg *types.MsgCraeteValue) (*types.MsgCraeteValueResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCraeteValueResponse{}, nil
}
