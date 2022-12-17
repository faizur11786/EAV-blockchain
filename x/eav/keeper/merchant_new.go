package keeper

import (
	"belShare/x/eav/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetMerchantNew set a specific merchantNew in the store from its index
func (k Keeper) SetMerchantNew(ctx sdk.Context, merchantNew types.MerchantNew) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MerchantNewKeyPrefix))
	b := k.cdc.MustMarshal(&merchantNew)
	store.Set(types.MerchantNewKey(
		merchantNew.Address,
	), b)
}

// GetMerchantNew returns a merchantNew from its index
func (k Keeper) GetMerchantNew(
	ctx sdk.Context,
	address string,

) (val types.MerchantNew, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MerchantNewKeyPrefix))

	b := store.Get(types.MerchantNewKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMerchantNew removes a merchantNew from the store
func (k Keeper) RemoveMerchantNew(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MerchantNewKeyPrefix))
	store.Delete(types.MerchantNewKey(
		address,
	))
}

// GetAllMerchantNew returns all merchantNew
func (k Keeper) GetAllMerchantNew(ctx sdk.Context) (list []types.MerchantNew) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MerchantNewKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.MerchantNew
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
