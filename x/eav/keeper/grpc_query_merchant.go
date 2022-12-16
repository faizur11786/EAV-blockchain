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

func (k Keeper) MerchantAll(c context.Context, req *types.QueryAllMerchantRequest) (*types.QueryAllMerchantResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var merchants []types.Merchant
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	merchantStore := prefix.NewStore(store, types.KeyPrefix(types.MerchantKeyPrefix))

	pageRes, err := query.Paginate(merchantStore, req.Pagination, func(key []byte, value []byte) error {
		var merchant types.Merchant
		if err := k.cdc.Unmarshal(value, &merchant); err != nil {
			return err
		}

		merchants = append(merchants, merchant)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMerchantResponse{Merchant: merchants, Pagination: pageRes}, nil
}

func (k Keeper) Merchant(c context.Context, req *types.QueryGetMerchantRequest) (*types.QueryGetMerchantResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetMerchant(
		ctx,
		req.Guid,
		req.Creator,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetMerchantResponse{Merchant: val}, nil
}
