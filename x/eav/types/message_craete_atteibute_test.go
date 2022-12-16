package types

import (
	"testing"

	"belShare/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCraeteAtteibute_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCraeteAtteibute
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCraeteAtteibute{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCraeteAtteibute{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
