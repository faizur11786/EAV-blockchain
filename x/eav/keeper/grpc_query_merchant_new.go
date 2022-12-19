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

func (k Keeper) MerchantNewAll(c context.Context, req *types.QueryAllMerchantNewRequest) (*types.QueryAllMerchantNewResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var merchantNews []types.MerchantNew
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	merchantNewStore := prefix.NewStore(store, types.KeyPrefix(types.MerchantNewKeyPrefix))

	pageRes, err := query.Paginate(merchantNewStore, req.Pagination, func(key []byte, value []byte) error {
		var merchantNew types.MerchantNew
		if err := k.cdc.Unmarshal(value, &merchantNew); err != nil {
			return err
		}

		merchantNews = append(merchantNews, merchantNew)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMerchantNewResponse{MerchantNew: merchantNews, Pagination: pageRes}, nil
}

func (k Keeper) MerchantNew(c context.Context, req *types.QueryGetMerchantNewRequest) (*types.QueryGetMerchantNewResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	_, found := k.GetEntityType(ctx, req.EntityId)
	if !found {
		return nil, status.Error(codes.NotFound, "Entity not found",)
	}
	
	val, found := k.GetMerchantNew(
		ctx,
		req.Address,
	)
	
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}
	
	attributes := k.GetAttributesByEntity(ctx, req.EntityId)


	var values []types.Value

	for _, attribute := range attributes{
		value, _ := k.GetValue(ctx, val.Guid, attribute.Guid)
		values = append(values, value)
	}


	return &types.QueryGetMerchantNewResponse{MerchantNew: val, Attributes: attributes, Values: values}, nil
}
