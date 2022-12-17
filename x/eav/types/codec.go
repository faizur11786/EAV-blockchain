package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateUser{}, "eav/CreateUser", nil)
	cdc.RegisterConcrete(&MsgCreateEntityType{}, "eav/CreateEntityType", nil)
	cdc.RegisterConcrete(&MsgCraeteAtteibute{}, "eav/CraeteAtteibute", nil)
	cdc.RegisterConcrete(&MsgCraeteValue{}, "eav/CraeteValue", nil)

	cdc.RegisterConcrete(&MsgCreateShop{}, "eav/CreateShop", nil)
	cdc.RegisterConcrete(&MsgNewMerchant{}, "eav/NewMerchant", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateUser{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateEntityType{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCraeteAtteibute{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCraeteValue{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateShop{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgNewMerchant{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
