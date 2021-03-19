package registry

type listMessageOptions struct {
	exactFullName *string
}

func LMOWithFullName(name string) ListMessageOption {
	return func(options listMessageOptions) listMessageOptions {
		options.exactFullName = &name
		return options
	}
}
