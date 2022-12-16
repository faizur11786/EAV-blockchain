package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCraeteAtteibute = "craete_atteibute"

var _ sdk.Msg = &MsgCraeteAtteibute{}

func NewMsgCraeteAtteibute(creator string, name string, entityId string) *MsgCraeteAtteibute {
	return &MsgCraeteAtteibute{
		Creator:  creator,
		Name:     name,
		EntityId: entityId,
	}
}

func (msg *MsgCraeteAtteibute) Route() string {
	return RouterKey
}

func (msg *MsgCraeteAtteibute) Type() string {
	return TypeMsgCraeteAtteibute
}

func (msg *MsgCraeteAtteibute) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCraeteAtteibute) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCraeteAtteibute) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
