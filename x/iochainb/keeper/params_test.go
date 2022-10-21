package keeper_test

import (
	"testing"

	testkeeper "github.com/mytestlab123/iochainb/testutil/keeper"
	"github.com/mytestlab123/iochainb/x/iochainb/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IochainbKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
