package keeper

import (
	"context"
	"fmt"

	"belShare/x/eav/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

func (k msgServer) CraeteAtteibute(goCtx context.Context, msg *types.MsgCraeteAtteibute) (*types.MsgCraeteAtteibuteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the entity in exsits
	_, found := k.GetEntityType(ctx, msg.EntityId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, fmt.Sprintf("Entity %d doesn't exist", msg.EntityId))
	}

	//  Create variable of atteribute
	id := uuid.New().String()
	var atteribute = types.Attribute{
		Guid:     id,
		Name:     msg.Name,
		EntityId: msg.EntityId,
	}
	// Write Attribute information to the store
	k.SetAttribute(ctx, atteribute)

	return &types.MsgCraeteAtteibuteResponse{Id: id}, nil
}
