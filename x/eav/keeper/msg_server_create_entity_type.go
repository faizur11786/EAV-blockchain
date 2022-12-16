package keeper

import (
	"context"

	"belShare/x/eav/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
)

func (k msgServer) CreateEntityType(goCtx context.Context, msg *types.MsgCreateEntityType) (*types.MsgCreateEntityTypeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := uuid.New().String()
	// Create variable of type entity type
	var entityType = types.EntityType{
		Guid:    id,
		Name:    msg.Name,
		Creator: msg.GetCreator(),
	}

	// Write user information to the store
	k.SetEntityType(ctx, entityType)

	return &types.MsgCreateEntityTypeResponse{Id: id}, nil
}
