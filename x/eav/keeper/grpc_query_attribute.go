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

func (k Keeper) AttributeAll(c context.Context, req *types.QueryAllAttributeRequest) (*types.QueryAllAttributeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var attributes []types.Attribute
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	attributeStore := prefix.NewStore(store, types.KeyPrefix(types.AttributeKeyPrefix))

	pageRes, err := query.Paginate(attributeStore, req.Pagination, func(key []byte, value []byte) error {
		var attribute types.Attribute
		if err := k.cdc.Unmarshal(value, &attribute); err != nil {
			return err
		}

		attributes = append(attributes, attribute)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAttributeResponse{Attribute: attributes, Pagination: pageRes}, nil
}

func (k Keeper) Attribute(c context.Context, req *types.QueryGetAttributeRequest) (*types.QueryGetAttributeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetAttribute(
		ctx,
		req.Guid,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetAttributeResponse{Attribute: val}, nil
}
