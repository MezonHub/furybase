package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreatePool{}, "fdex/CreatePool", nil)
	cdc.RegisterConcrete(&MsgAddLiquidity{}, "fdex/AddLiquidity", nil)
	cdc.RegisterConcrete(&MsgSwap{}, "fdex/Swap", nil)
	cdc.RegisterConcrete(&MsgRemoveLiquidity{}, "fdex/RemoveLiquidity", nil)
	cdc.RegisterConcrete(&MsgToggleProviderSwitch{}, "fdex/ToggleProviderSwitch", nil)
	cdc.RegisterConcrete(&MsgAddProvider{}, "fdex/AddProvider", nil)
	cdc.RegisterConcrete(&MsgRmProvider{}, "fdex/RmProvider", nil)
	cdc.RegisterConcrete(&MsgAddPoolCreator{}, "fdex/AddPoolCreator", nil)
	cdc.RegisterConcrete(&MsgRmPoolCreator{}, "fdex/RmPoolCreator", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddLiquidity{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSwap{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveLiquidity{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgToggleProviderSwitch{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddProvider{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRmProvider{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddPoolCreator{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRmPoolCreator{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(Amino)
)

func init() {
	RegisterCodec(Amino)
	cryptocodec.RegisterCrypto(Amino)
	sdk.RegisterLegacyAminoCodec(Amino)
	Amino.Seal()
}
