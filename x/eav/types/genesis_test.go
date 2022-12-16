package types_test

import (
	"testing"

	"belShare/x/eav/types"

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
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				MerchantList: []types.Merchant{
					{
						Guid:    "0",
						Creator: "0",
					},
					{
						Guid:    "1",
						Creator: "1",
					},
				},
				EntityTypeList: []types.EntityType{
					{
						Guid: "0",
					},
					{
						Guid: "1",
					},
				},
				AttributeList: []types.Attribute{
					{
						Guid: "0",
					},
					{
						Guid: "1",
					},
				},
				ValueList: []types.Value{
					{
						EntityId:    "0",
						AttributeId: "0",
					},
					{
						EntityId:    "1",
						AttributeId: "1",
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
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated merchant",
			genState: &types.GenesisState{
				MerchantList: []types.Merchant{
					{
						Guid:    "0",
						Creator: "0",
					},
					{
						Guid:    "0",
						Creator: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated entityType",
			genState: &types.GenesisState{
				EntityTypeList: []types.EntityType{
					{
						Guid: "0",
					},
					{
						Guid: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated attribute",
			genState: &types.GenesisState{
				AttributeList: []types.Attribute{
					{
						Guid: "0",
					},
					{
						Guid: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated value",
			genState: &types.GenesisState{
				ValueList: []types.Value{
					{
						EntityId:    "0",
						AttributeId: "0",
					},
					{
						EntityId:    "0",
						AttributeId: "0",
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
