package keeper

import (
	"context"
	"fmt"

	"belShare/x/eav/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

func (k msgServer) CreateShop(goCtx context.Context, msg *types.MsgCreateShop) (*types.MsgCreateShopResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	id := uuid.New()
	for _, result := range msg.Attributs{
		attribute, found := k.GetAttribute(ctx, result.Guid)
		if !found {
			return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, fmt.Sprintf("Attribute %d doesn't exist", result.Guid))
		}
		k.SetValue(ctx, types.Value{
			EntityId: id.String(),
			AttributeId: attribute.Guid,
			Guid: uuid.New().String(),
			Value: result.Name,
		})
	}

	// Create variable of type Shop
	var shop = types.Merchant{
		Guid: id.String(),
		Creator: msg.Creator,
		Address: msg.Address,
	}

	k.SetMerchant(ctx, shop)

	return &types.MsgCreateShopResponse{}, nil
}
