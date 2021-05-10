package registry

type ListMessageOptions struct {
	ExactFullName *string
	PrefixMatch   *string
}

// does message sub definition based resolution
func LMOWithFullName(name string) ListMessageOption {
	return func(options ListMessageOptions) ListMessageOptions {
		options.ExactFullName = &name
		return options
	}
}
// arg: prefix is usually longer than what is matched
// used for field name resolution;
func LMOWithPrefixMatch(prefix string) ListMessageOption {
	return func(options ListMessageOptions) ListMessageOptions {
		options.PrefixMatch = &prefix
		return options
	}
}
