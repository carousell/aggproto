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
