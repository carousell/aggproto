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
	"encoding/json"
	"log"
	"net/http"
)

func responseHandler(writer http.ResponseWriter, resp interface{}, err error) {
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(err.Error()))
	}
	bytes, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(err.Error()))
	}
	_, _ = writer.Write(bytes)
}
