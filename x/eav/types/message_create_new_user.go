package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateNewUser = "create_new_user"

var _ sdk.Msg = &MsgCreateNewUser{}

func NewMsgCreateNewUser(creator string, address string, attributes []*Attribute) *MsgCreateNewUser {
	return &MsgCreateNewUser{
		Creator:    creator,
		Address:    address,
		Attributes: attributes,
	}
}

func (msg *MsgCreateNewUser) Route() string {
	return RouterKey
}

func (msg *MsgCreateNewUser) Type() string {
	return TypeMsgCreateNewUser
}

func (msg *MsgCreateNewUser) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateNewUser) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateNewUser) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
