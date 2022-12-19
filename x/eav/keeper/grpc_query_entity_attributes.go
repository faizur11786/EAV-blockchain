package keeper

import (
	"context"

	"belShare/x/eav/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EntityAttributes(goCtx context.Context, req *types.QueryEntityAttributesRequest) (*types.QueryEntityAttributesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetEntityType(ctx, req.EntityId)
	if !found {
		return nil, status.Error(codes.NotFound, "Entity not found")
	}

	entities := k.GetAllAttribute(ctx)

	var entityAtt []*types.Attribute

	for _, entity := range entities {
		if entity.EntityId == req.EntityId {
			entityAtt = append(entityAtt, &types.Attribute{
				Guid:     entity.Guid,
				Name:     entity.Name,
				EntityId: entity.EntityId,
			})
		}
	}

	return &types.QueryEntityAttributesResponse{Attributes: entityAtt}, nil
}

func (k Keeper) GetAttributesByEntity(ctx sdk.Context, entityId string) (entityAtt []types.Attribute) {

	entities := k.GetAllAttribute(ctx)

	for _, entity := range entities {
		if entity.EntityId == entityId {
			entityAtt = append(entityAtt, types.Attribute{
				Guid:     entity.Guid,
				Name:     entity.Name,
				EntityId: entity.EntityId,
			})
		}
	}

	return
}
