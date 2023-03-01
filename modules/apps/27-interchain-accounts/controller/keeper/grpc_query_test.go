package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/treasurenetprotocol/treasurenet_ibc/modules/apps/27-interchain-accounts/controller/types"
)

func (suite *KeeperTestSuite) TestQueryParams() {
	ctx := sdk.WrapSDKContext(suite.chainA.GetContext())
	expParams := types.DefaultParams()
	res, _ := suite.chainA.GetSimApp().ICAControllerKeeper.Params(ctx, &types.QueryParamsRequest{})
	suite.Require().Equal(&expParams, res.Params)
}
