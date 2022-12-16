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

func TestMerchantQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.EavKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNMerchant(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetMerchantRequest
		response *types.QueryGetMerchantResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetMerchantRequest{
				Guid:    msgs[0].Guid,
				Creator: msgs[0].Creator,
			},
			response: &types.QueryGetMerchantResponse{Merchant: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetMerchantRequest{
				Guid:    msgs[1].Guid,
				Creator: msgs[1].Creator,
			},
			response: &types.QueryGetMerchantResponse{Merchant: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetMerchantRequest{
				Guid:    strconv.Itoa(100000),
				Creator: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Merchant(wctx, tc.request)
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

func TestMerchantQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.EavKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNMerchant(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllMerchantRequest {
		return &types.QueryAllMerchantRequest{
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
			resp, err := keeper.MerchantAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Merchant), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Merchant),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.MerchantAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Merchant), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Merchant),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.MerchantAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Merchant),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.MerchantAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
