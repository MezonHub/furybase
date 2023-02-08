package types

import (
// this line is used by starport scaffolding # genesis/types/import
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:                 DefaultParams(),
		SelectedFValidatorList: []*FValidator{},
		LatestVotedCycleList:   []*Cycle{},
		LatestDealedCycleList:  []*Cycle{},
		CycleSecondsList:       []*CycleSeconds{},
		ShuffleSecondsList:     []*ShuffleSeconds{},
		DealingFValidatorList:  []*DealingFValidator{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
