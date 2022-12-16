package eav

import (
	"math/rand"

	"belShare/testutil/sample"
	eavsimulation "belShare/x/eav/simulation"
	"belShare/x/eav/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = eavsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateUser = "op_weight_msg_create_user"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateUser int = 100

	opWeightMsgCreateMerchant = "op_weight_msg_create_merchant"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMerchant int = 100

	opWeightMsgCreateEntityType = "op_weight_msg_create_entity_type"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateEntityType int = 100

	opWeightMsgCraeteAtteibute = "op_weight_msg_craete_atteibute"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCraeteAtteibute int = 100

	opWeightMsgCraeteValue = "op_weight_msg_craete_value"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCraeteValue int = 100

	opWeightMsgCreateShop = "op_weight_msg_create_shop"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateShop int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	eavGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&eavGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateUser int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateUser, &weightMsgCreateUser, nil,
		func(_ *rand.Rand) {
			weightMsgCreateUser = defaultWeightMsgCreateUser
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateUser,
		eavsimulation.SimulateMsgCreateUser(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateEntityType int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateEntityType, &weightMsgCreateEntityType, nil,
		func(_ *rand.Rand) {
			weightMsgCreateEntityType = defaultWeightMsgCreateEntityType
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateEntityType,
		eavsimulation.SimulateMsgCreateEntityType(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCraeteAtteibute int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCraeteAtteibute, &weightMsgCraeteAtteibute, nil,
		func(_ *rand.Rand) {
			weightMsgCraeteAtteibute = defaultWeightMsgCraeteAtteibute
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCraeteAtteibute,
		eavsimulation.SimulateMsgCraeteAtteibute(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCraeteValue int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCraeteValue, &weightMsgCraeteValue, nil,
		func(_ *rand.Rand) {
			weightMsgCraeteValue = defaultWeightMsgCraeteValue
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCraeteValue,
		eavsimulation.SimulateMsgCraeteValue(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateShop int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateShop, &weightMsgCreateShop, nil,
		func(_ *rand.Rand) {
			weightMsgCreateShop = defaultWeightMsgCreateShop
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateShop,
		eavsimulation.SimulateMsgCreateShop(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
