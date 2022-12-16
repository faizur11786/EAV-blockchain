package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "belShare/testutil/keeper"
	"belShare/testutil/nullify"
	"belShare/x/eav/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestAttributeQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.EavKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAttribute(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetAttributeRequest
		response *types.QueryGetAttributeResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetAttributeRequest{
				Guid: msgs[0].Guid,
			},
			response: &types.QueryGetAttributeResponse{Attribute: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetAttributeRequest{
				Guid: msgs[1].Guid,
			},
			response: &types.QueryGetAttributeResponse{Attribute: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetAttributeRequest{
				Guid: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Attribute(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestAttributeQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.EavKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAttribute(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllAttributeRequest {
		return &types.QueryAllAttributeRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.AttributeAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Attribute), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Attribute),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.AttributeAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Attribute), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Attribute),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.AttributeAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Attribute),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.AttributeAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
