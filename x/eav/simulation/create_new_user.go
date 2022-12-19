package simulation

import (
	"math/rand"

	"belShare/x/eav/keeper"
	"belShare/x/eav/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCreateNewUser(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateNewUser{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateNewUser simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateNewUser simulation not implemented"), nil, nil
	}
}
