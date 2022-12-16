package keeper

import (
	"belShare/x/eav/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetEntityType set a specific entityType in the store from its index
func (k Keeper) SetEntityType(ctx sdk.Context, entityType types.EntityType) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EntityTypeKeyPrefix))
	b := k.cdc.MustMarshal(&entityType)
	store.Set(types.EntityTypeKey(
		entityType.Guid,
	), b)
}

// GetEntityType returns a entityType from its index
func (k Keeper) GetEntityType(
	ctx sdk.Context,
	guid string,

) (val types.EntityType, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EntityTypeKeyPrefix))

	b := store.Get(types.EntityTypeKey(
		guid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEntityType removes a entityType from the store
func (k Keeper) RemoveEntityType(
	ctx sdk.Context,
	guid string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EntityTypeKeyPrefix))
	store.Delete(types.EntityTypeKey(
		guid,
	))
}

// GetAllEntityType returns all entityType
func (k Keeper) GetAllEntityType(ctx sdk.Context) (list []types.EntityType) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EntityTypeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.EntityType
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
