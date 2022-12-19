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

func (k Keeper) NewUserAll(c context.Context, req *types.QueryAllNewUserRequest) (*types.QueryAllNewUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var newUsers []types.NewUser
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	newUserStore := prefix.NewStore(store, types.KeyPrefix(types.NewUserKeyPrefix))

	pageRes, err := query.Paginate(newUserStore, req.Pagination, func(key []byte, value []byte) error {
		var newUser types.NewUser
		if err := k.cdc.Unmarshal(value, &newUser); err != nil {
			return err
		}

		newUsers = append(newUsers, newUser)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNewUserResponse{NewUser: newUsers, Pagination: pageRes}, nil
}

func (k Keeper) NewUser(c context.Context, req *types.QueryGetNewUserRequest) (*types.QueryGetNewUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNewUser(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNewUserResponse{NewUser: val}, nil
}
