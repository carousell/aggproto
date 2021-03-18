package opresolution

import "github.com/carousell/aggproto/pkg/generator/stages/msgresolution"

type OperationResolver interface {
	Resolve(ctx msgresolution.AdaptorContext)
}

type OperationContext interface {

}
