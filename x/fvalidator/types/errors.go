package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/fvalidator module sentinel errors
var (
	ErrFValidatorAlreadyExist              = sdkerrors.Register(ModuleName, 1101, "fValidator already exist")
	ErrFValidatorNotExist                  = sdkerrors.Register(ModuleName, 1102, "fValidator not exist")
	ErrCycleBehindLatestCycle              = sdkerrors.Register(ModuleName, 1103, "cycle behind latest voted cycle")
	ErrCycleVersionNotMatch                = sdkerrors.Register(ModuleName, 1104, "cycle version not match")
	ErrLatestVotedCycleNotDealed           = sdkerrors.Register(ModuleName, 1105, "latest voted cycle not dealed")
	ErrLedgerIsBusyWithEra                 = sdkerrors.Register(ModuleName, 1106, "ledger is busy with era")
	ErrReportCycleNotMatchLatestVotedCycle = sdkerrors.Register(ModuleName, 1107, "report cycle not match latest voted cycle")
	ErrLedgerChainEraNotExist              = sdkerrors.Register(ModuleName, 1108, "ledger chain era not exist")
	ErrDealingFvalidatorNotFound           = sdkerrors.Register(ModuleName, 1109, "dealing fvalidator not found")
	ErrOldEqualNewFValidator               = sdkerrors.Register(ModuleName, 1110, "old euqal new fValidator")
	ErrFValidatorAlreadyInit               = sdkerrors.Register(ModuleName, 1111, "fValidator already init")
)
