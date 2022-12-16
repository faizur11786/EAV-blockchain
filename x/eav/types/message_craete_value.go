package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCraeteValue = "craete_value"

var _ sdk.Msg = &MsgCraeteValue{}

func NewMsgCraeteValue(creator string, attributeId string, entityId string, value string) *MsgCraeteValue {
	return &MsgCraeteValue{
		Creator:     creator,
		AttributeId: attributeId,
		EntityId:    entityId,
		Value:       value,
	}
}

func (msg *MsgCraeteValue) Route() string {
	return RouterKey
}

func (msg *MsgCraeteValue) Type() string {
	return TypeMsgCraeteValue
}

func (msg *MsgCraeteValue) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCraeteValue) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCraeteValue) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
