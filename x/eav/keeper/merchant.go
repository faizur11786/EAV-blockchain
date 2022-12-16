package keeper

import (
	"belShare/x/eav/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetMerchant set a specific merchant in the store from its index
func (k Keeper) SetMerchant(ctx sdk.Context, merchant types.Merchant) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MerchantKeyPrefix))
	b := k.cdc.MustMarshal(&merchant)
	store.Set(types.MerchantKey(
		merchant.Guid,
		merchant.Creator,
	), b)
}

// GetMerchant returns a merchant from its index
func (k Keeper) GetMerchant(
	ctx sdk.Context,
	guid string,
	creator string,

) (val types.Merchant, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MerchantKeyPrefix))

	b := store.Get(types.MerchantKey(
		guid,
		creator,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMerchant removes a merchant from the store
func (k Keeper) RemoveMerchant(
	ctx sdk.Context,
	guid string,
	creator string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MerchantKeyPrefix))
	store.Delete(types.MerchantKey(
		guid,
		creator,
	))
}

// GetAllMerchant returns all merchant
func (k Keeper) GetAllMerchant(ctx sdk.Context) (list []types.Merchant) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MerchantKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Merchant
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
