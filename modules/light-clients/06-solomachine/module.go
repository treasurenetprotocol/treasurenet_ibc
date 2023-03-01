package solomachine

import (
	"github.com/treasurenetprotocol/treasurenet_ibc/modules/light-clients/06-solomachine/types"
)

// Name returns the solo machine client name.
func Name() string {
	return types.SubModuleName
}
