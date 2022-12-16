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

func TestEntityTypeQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.EavKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNEntityType(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetEntityTypeRequest
		response *types.QueryGetEntityTypeResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetEntityTypeRequest{
				Guid: msgs[0].Guid,
			},
			response: &types.QueryGetEntityTypeResponse{EntityType: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetEntityTypeRequest{
				Guid: msgs[1].Guid,
			},
			response: &types.QueryGetEntityTypeResponse{EntityType: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetEntityTypeRequest{
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
			response, err := keeper.EntityType(wctx, tc.request)
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

func TestEntityTypeQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.EavKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNEntityType(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllEntityTypeRequest {
		return &types.QueryAllEntityTypeRequest{
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
			resp, err := keeper.EntityTypeAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.EntityType), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.EntityType),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.EntityTypeAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.EntityType), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.EntityType),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.EntityTypeAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.EntityType),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.EntityTypeAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
