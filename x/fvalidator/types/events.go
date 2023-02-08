package types

// sudo module event types
const (
	EventTypeInitFValidator   = "init_fvalidator"
	EventTypeAddFValidator    = "add_fvalidator"
	EventTypeUpdateFValidator = "update_fvalidator"

	AttributeKeyDenom        = "denom"
	AttributeKeyAddresses    = "addresses"
	AttributeKeyAddress      = "address"
	AttributeKeyNewAddress   = "new_address"
	AttributeKeyOldAddress   = "old_address"
	AttributeKeyAddedAddress = "added_address"
	AttributeKeyPoolAddress  = "pool_address"
	AttributeKeyChainEra     = "chain_era"
	AttributeKeyCycleVersion = "cycle_version"
	AttributeKeyCycleNumber  = "cycle_number"
	AttributeKeyCycleSeconds = "cycle_seconds"
)
