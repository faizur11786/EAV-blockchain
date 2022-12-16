package keeper

import (
	"context"

	"belShare/x/eav/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EntityTypeAll(c context.Context, req *types.QueryAllEntityTypeRequest) (*types.QueryAllEntityTypeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var entityTypes []types.EntityType
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	entityTypeStore := prefix.NewStore(store, types.KeyPrefix(types.EntityTypeKeyPrefix))

	pageRes, err := query.Paginate(entityTypeStore, req.Pagination, func(key []byte, value []byte) error {
		var entityType types.EntityType
		if err := k.cdc.Unmarshal(value, &entityType); err != nil {
			return err
		}

		entityTypes = append(entityTypes, entityType)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEntityTypeResponse{EntityType: entityTypes, Pagination: pageRes}, nil
}

func (k Keeper) EntityType(c context.Context, req *types.QueryGetEntityTypeRequest) (*types.QueryGetEntityTypeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetEntityType(
		ctx,
		req.Guid,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetEntityTypeResponse{EntityType: val}, nil
}
