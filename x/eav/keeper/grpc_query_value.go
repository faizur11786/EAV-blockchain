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

func (k Keeper) ValueAll(c context.Context, req *types.QueryAllValueRequest) (*types.QueryAllValueResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var values []types.Value
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	valueStore := prefix.NewStore(store, types.KeyPrefix(types.ValueKeyPrefix))

	pageRes, err := query.Paginate(valueStore, req.Pagination, func(key []byte, dataValue []byte) error {
		var value types.Value
		if err := k.cdc.Unmarshal(dataValue, &value); err != nil {
			return err
		}

		values = append(values, value)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllValueResponse{Value: values, Pagination: pageRes}, nil
}

func (k Keeper) Value(c context.Context, req *types.QueryGetValueRequest) (*types.QueryGetValueResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetValue(
		ctx,
		req.EntityId,
		req.AttributeId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetValueResponse{Value: val}, nil
}
