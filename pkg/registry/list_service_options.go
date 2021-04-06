package registry

type ListServiceOptions struct {
	MatchReturn Message
	OpNameIn    []string
}

func LSOWithOutputMessage(msg Message) ListServiceOption {
	return func(options ListServiceOptions) ListServiceOptions {
		options.MatchReturn = msg
		return options
	}
}

func LSOWithOperationNameIn(names []string) ListServiceOption {
	return func(options ListServiceOptions) ListServiceOptions {
		options.OpNameIn = names
		return options
	}
}
