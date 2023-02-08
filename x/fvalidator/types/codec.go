package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	fvotetypes "github.com/furybase/furybase/x/fvote/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgInitFValidator{}, "fvalidator/InitFValidator", nil)
	cdc.RegisterConcrete(&UpdateFValidatorProposal{}, "fvalidator/UpdateFValidator", nil)
	cdc.RegisterConcrete(&UpdateFValidatorReportProposal{}, "fvalidator/UpdateFValidatorReport", nil)
	cdc.RegisterConcrete(&MsgSetCycleSeconds{}, "fvalidator/SetCycleSeconds", nil)
	cdc.RegisterConcrete(&MsgSetShuffleSeconds{}, "fvalidator/SetShuffleSeconds", nil)
	cdc.RegisterConcrete(&MsgAddFValidator{}, "fvalidator/AddFValidator", nil)
	cdc.RegisterConcrete(&MsgRmFValidator{}, "fvalidator/RmFValidator", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInitFValidator{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&UpdateFValidatorProposal{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetCycleSeconds{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetShuffleSeconds{},
	)
	registry.RegisterImplementations(
		(*fvotetypes.Content)(nil),
		&UpdateFValidatorProposal{},
	)
	registry.RegisterImplementations(
		(*fvotetypes.Content)(nil),
		&UpdateFValidatorReportProposal{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddFValidator{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRmFValidator{},
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
