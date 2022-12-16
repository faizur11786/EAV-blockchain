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

func TestValueQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.EavKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNValue(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetValueRequest
		response *types.QueryGetValueResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetValueRequest{
				EntityId:    msgs[0].EntityId,
				AttributeId: msgs[0].AttributeId,
			},
			response: &types.QueryGetValueResponse{Value: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetValueRequest{
				EntityId:    msgs[1].EntityId,
				AttributeId: msgs[1].AttributeId,
			},
			response: &types.QueryGetValueResponse{Value: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetValueRequest{
				EntityId:    strconv.Itoa(100000),
				AttributeId: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Value(wctx, tc.request)
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

func TestValueQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.EavKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNValue(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllValueRequest {
		return &types.QueryAllValueRequest{
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
			resp, err := keeper.ValueAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Value), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Value),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ValueAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Value), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Value),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ValueAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Value),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ValueAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
