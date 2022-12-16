package types_test

import (
	"testing"

	"belShare/x/user/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				UserList: []types.User{
					{
						Creator: "0",
						Guid:    "0",
					},
					{
						Creator: "1",
						Guid:    "1",
					},
				},
				XxxList: []types.Xxx{
					{
						Address: "0",
						Creator: "0",
					},
					{
						Address: "1",
						Creator: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated user",
			genState: &types.GenesisState{
				UserList: []types.User{
					{
						Creator: "0",
						Guid:    "0",
					},
					{
						Creator: "0",
						Guid:    "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated xxx",
			genState: &types.GenesisState{
				XxxList: []types.Xxx{
					{
						Address: "0",
						Creator: "0",
					},
					{
						Address: "0",
						Creator: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
