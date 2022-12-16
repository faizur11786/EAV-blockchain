package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateShop = "create_shop"

var _ sdk.Msg = &MsgCreateShop{}

func NewMsgCreateShop(creator string, address string, attributs []*Attribute) *MsgCreateShop {
	return &MsgCreateShop{
		Creator:   creator,
		Address:   address,
		Attributs: attributs,
	}
}

func (msg *MsgCreateShop) Route() string {
	return RouterKey
}

func (msg *MsgCreateShop) Type() string {
	return TypeMsgCreateShop
}

func (msg *MsgCreateShop) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateShop) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateShop) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
