package registry

type listMessageOptions struct {
	exactFullName *string
	prefixMatch   *string
}

func LMOWithFullName(name string) ListMessageOption {
	return func(options listMessageOptions) listMessageOptions {
		options.exactFullName = &name
		return options
	}
}

func LMOWithPrefixMatch(prefix string) ListMessageOption {
	return func(options listMessageOptions) listMessageOptions {
		options.prefixMatch = &prefix
		return options
	}
}
