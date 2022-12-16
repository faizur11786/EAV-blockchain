package keeper

import (
	"belShare/x/eav/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetAttribute set a specific attribute in the store from its index
func (k Keeper) SetAttribute(ctx sdk.Context, attribute types.Attribute) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AttributeKeyPrefix))
	b := k.cdc.MustMarshal(&attribute)
	store.Set(types.AttributeKey(
		attribute.Guid,
	), b)
}

// GetAttribute returns a attribute from its index
func (k Keeper) GetAttribute(
	ctx sdk.Context,
	guid string,

) (val types.Attribute, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AttributeKeyPrefix))

	b := store.Get(types.AttributeKey(
		guid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAttribute removes a attribute from the store
func (k Keeper) RemoveAttribute(
	ctx sdk.Context,
	guid string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AttributeKeyPrefix))
	store.Delete(types.AttributeKey(
		guid,
	))
}

// GetAllAttribute returns all attribute
func (k Keeper) GetAllAttribute(ctx sdk.Context) (list []types.Attribute) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AttributeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Attribute
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
