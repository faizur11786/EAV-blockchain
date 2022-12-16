package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateEntityType = "create_entity_type"

var _ sdk.Msg = &MsgCreateEntityType{}

func NewMsgCreateEntityType(creator string, name string) *MsgCreateEntityType {
	return &MsgCreateEntityType{
		Creator: creator,
		Name:    name,
	}
}

func (msg *MsgCreateEntityType) Route() string {
	return RouterKey
}

func (msg *MsgCreateEntityType) Type() string {
	return TypeMsgCreateEntityType
}

func (msg *MsgCreateEntityType) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateEntityType) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateEntityType) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
