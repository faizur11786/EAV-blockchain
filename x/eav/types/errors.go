package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/eav module sentinel errors
var (
	ErrSample      = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrDuplicateID = sdkerrors.Register(ModuleName, 1400, "")
)
