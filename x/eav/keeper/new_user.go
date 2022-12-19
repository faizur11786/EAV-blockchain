package keeper

import (
	"belShare/x/eav/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetNewUser set a specific newUser in the store from its index
func (k Keeper) SetNewUser(ctx sdk.Context, newUser types.NewUser) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NewUserKeyPrefix))
	b := k.cdc.MustMarshal(&newUser)
	store.Set(types.NewUserKey(
		newUser.Address,
	), b)
}

// GetNewUser returns a newUser from its index
func (k Keeper) GetNewUser(
	ctx sdk.Context,
	address string,

) (val types.NewUser, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NewUserKeyPrefix))

	b := store.Get(types.NewUserKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNewUser removes a newUser from the store
func (k Keeper) RemoveNewUser(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NewUserKeyPrefix))
	store.Delete(types.NewUserKey(
		address,
	))
}

// GetAllNewUser returns all newUser
func (k Keeper) GetAllNewUser(ctx sdk.Context) (list []types.NewUser) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NewUserKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NewUser
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
