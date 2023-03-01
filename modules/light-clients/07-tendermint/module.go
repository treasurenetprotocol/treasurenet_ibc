package tendermint

import (
	"github.com/treasurenetprotocol/treasurenet_ibc/modules/light-clients/07-tendermint/types"
)

// Name returns the IBC client name
func Name() string {
	return types.SubModuleName
}
