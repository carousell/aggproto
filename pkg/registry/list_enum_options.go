package registry

type ListEnumOptions struct {
	ExactFullName *string
}

func LEOWithFullName(name string) ListEnumOption {
	return func(options ListEnumOptions) ListEnumOptions {
		options.ExactFullName = &name
		return options
	}
}
