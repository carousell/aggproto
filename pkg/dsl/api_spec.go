package dsl

type ApiSpec struct {
	Api struct {
		Name    string
		Group   string
		Version int
	}
	Meta struct {
		GoPackage string
	}
	Input      []string
	Output     []string
	Operations []string
}

func (s *ApiSpec) ToContext() (*Context, error) {
	argDesc, err := parseArgDescriptors(s.Input...)
	if err != nil {
		return nil, err
	}
	return &Context{
		Api: &apiDescriptor{
			name:    s.Api.Name,
			group:   s.Api.Group,
			version: s.Api.Version,
		},
		Operation: OpDescriptor{AllowedOperations: s.Operations},
		Output:    parseOutputFields(s.Output...),
		Meta:      Meta{GoPackage: s.Meta.GoPackage},
		Input:     argDesc,
	}, nil
}
