//
//  This source file is part of the carousell/aggproto open source project
//
//  Copyright © 2021 Carousell and the project authors
//  Licensed under Apache License v2.0
//
//  See https://github.com/carousell/aggproto/blob/master/LICENSE for license information
//  See https://github.com/carousell/aggproto/graphs/contributors for the list of project authors
//
package apis

import (
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	spec *apiSpecs
	reg  *registry
}

func (s *server) RegisterHttp() {
	r := mux.NewRouter()
	r.HandleFunc("/api/specs/list", func(writer http.ResponseWriter, request *http.Request) {
		apiDescs, err := s.spec.list()
		responseHandler(writer, apiDescs, err)
	})
	r.HandleFunc("/api/specs/{specFile}", func(writer http.ResponseWriter, request *http.Request) {
		specFile := mux.Vars(request)["specFile"]
		spec, err := s.spec.load(specFile)
		responseHandler(writer, spec, err)
	})
	r.HandleFunc("/api/registry/search/message/{prefixString}", func(writer http.ResponseWriter, request *http.Request) {
		prefixString := mux.Vars(request)["prefixString"]
		if prefixString == "*" {
			prefixString = ""
		}
		msgs := s.reg.searchMessage(prefixString)
		responseHandler(writer, msgs, nil)
	})
	http.Handle("/api/", r)
}

func Server(specDir string, registryDir string) *server {
	reg := &registry{registryDir: registryDir}
	reg.init()
	return &server{spec: &apiSpecs{specDir: specDir}, reg: reg}
}
