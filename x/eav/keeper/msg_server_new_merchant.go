package keeper

import (
	"context"
	"fmt"

	"belShare/x/eav/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

func (k msgServer) NewMerchant(goCtx context.Context, msg *types.MsgNewMerchant) (*types.MsgNewMerchantResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)


	//  Check if merchain already exist
	_, found := k.GetMerchantNew(ctx, msg.Address)
	if found {
		return nil, sdkerrors.Wrap(types.ErrDuplicate, fmt.Sprintf("Merchent %d exist", msg.Address))
	}

	// Generate new merhcant's GUID
	id := uuid.New()
	// Create variable of type MerchantNew
	var shop = types.MerchantNew{
		Guid:    id.String(),
		Creator: msg.Creator,
		Address: msg.Address,
	}

	// Write Merchant information to the store 
	k.SetMerchantNew(ctx, shop)

	
	for _, result := range msg.Attributes {

		// Check if Attribute exist which merhcant want to store.
		attribute, found := k.GetAttribute(ctx, result.Guid)
		if !found {
			return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, fmt.Sprintf("Attribute %d doesn't exist", result.Guid))
		}

		// Check if Attribute-Value already exist 
		value, found := k.GetValue(ctx, id.String(), attribute.Guid);
		if found && value.Value == result.Name {
			return nil, sdkerrors.Wrap(types.ErrDuplicate, fmt.Sprintf("Value %d exist", result.Name))
		}

		// Write Value information to the store 
		k.SetValue(ctx, types.Value{
			EntityId:    id.String(),
			AttributeId: attribute.Guid,
			Guid:        uuid.New().String(),
			Value:       result.Name,
		})
	}



	return &types.MsgNewMerchantResponse{}, nil
}
