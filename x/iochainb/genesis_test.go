package iochainb_test

import (
	"testing"

	keepertest "github.com/mytestlab123/iochainb/testutil/keeper"
	"github.com/mytestlab123/iochainb/testutil/nullify"
	"github.com/mytestlab123/iochainb/x/iochainb"
	"github.com/mytestlab123/iochainb/x/iochainb/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IochainbKeeper(t)
	iochainb.InitGenesis(ctx, *k, genesisState)
	got := iochainb.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
