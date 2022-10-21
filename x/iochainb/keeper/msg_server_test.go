package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/mytestlab123/iochainb/testutil/keeper"
	"github.com/mytestlab123/iochainb/x/iochainb/keeper"
	"github.com/mytestlab123/iochainb/x/iochainb/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.IochainbKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
