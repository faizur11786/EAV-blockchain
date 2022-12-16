package keeper

import (
	"belShare/x/eav/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetValue set a specific value in the store from its index
func (k Keeper) SetValue(ctx sdk.Context, value types.Value) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValueKeyPrefix))
	b := k.cdc.MustMarshal(&value)
	store.Set(types.ValueKey(
		value.EntityId,
		value.AttributeId,
	), b)
}

// GetValue returns a value from its index
func (k Keeper) GetValue(
	ctx sdk.Context,
	entityId string,
	attributeId string,

) (val types.Value, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValueKeyPrefix))

	b := store.Get(types.ValueKey(
		entityId,
		attributeId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveValue removes a value from the store
func (k Keeper) RemoveValue(
	ctx sdk.Context,
	entityId string,
	attributeId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValueKeyPrefix))
	store.Delete(types.ValueKey(
		entityId,
		attributeId,
	))
}

// GetAllValue returns all value
func (k Keeper) GetAllValue(ctx sdk.Context) (list []types.Value) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValueKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Value
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
