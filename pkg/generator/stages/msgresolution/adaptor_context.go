package msgresolution

import (
	"github.com/carousell/aggproto/pkg/registry"
)

type adaptorContext struct {
	adaptorUnits []adaptorUnit
}

func (a *adaptorContext) Dependencies() []registry.Message {
	return nil
}
