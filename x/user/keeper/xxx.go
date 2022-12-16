package keeper

import (
	"belShare/x/user/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetXxx set a specific xxx in the store from its index
func (k Keeper) SetXxx(ctx sdk.Context, xxx types.Xxx) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.XxxKeyPrefix))
	b := k.cdc.MustMarshal(&xxx)
	store.Set(types.XxxKey(
		xxx.Address,
		xxx.Creator,
	), b)
}

// GetXxx returns a xxx from its index
func (k Keeper) GetXxx(
	ctx sdk.Context,
	address string,
	creator string,

) (val types.Xxx, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.XxxKeyPrefix))

	b := store.Get(types.XxxKey(
		address,
		creator,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveXxx removes a xxx from the store
func (k Keeper) RemoveXxx(
	ctx sdk.Context,
	address string,
	creator string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.XxxKeyPrefix))
	store.Delete(types.XxxKey(
		address,
		creator,
	))
}

// GetAllXxx returns all xxx
func (k Keeper) GetAllXxx(ctx sdk.Context) (list []types.Xxx) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.XxxKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Xxx
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
