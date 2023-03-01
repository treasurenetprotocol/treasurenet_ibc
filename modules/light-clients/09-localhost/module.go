package localhost

import (
	"github.com/treasurenetprotocol/treasurenet_ibc/modules/light-clients/09-localhost/types"
)

// Name returns the IBC client name
func Name() string {
	return types.SubModuleName
}
