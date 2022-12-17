package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgNewMerchant = "new_merchant"

var _ sdk.Msg = &MsgNewMerchant{}

func NewMsgNewMerchant(creator string, address string, attributes []*Attribute) *MsgNewMerchant {
	return &MsgNewMerchant{
		Creator:    creator,
		Address:    address,
		Attributes: attributes,
	}
}

func (msg *MsgNewMerchant) Route() string {
	return RouterKey
}

func (msg *MsgNewMerchant) Type() string {
	return TypeMsgNewMerchant
}

func (msg *MsgNewMerchant) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgNewMerchant) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgNewMerchant) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
