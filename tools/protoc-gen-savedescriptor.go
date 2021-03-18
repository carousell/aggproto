package main

import (
	"io/ioutil"
	"os"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	req := &pluginpb.CodeGeneratorRequest{}
	resp := &pluginpb.CodeGeneratorResponse{}
	err = proto.Unmarshal(data, req)
	if err != nil {
		panic(err)
	}
	parameters := req.GetParameter()
	for _, group := range strings.Split(parameters, ",") {
		kv := strings.Split(group, "=")
		if len(kv) > 1 {
		}

	}
	marshalled, err := proto.Marshal(resp)
	if err != nil {
		panic(err)
	}
	_, _ = os.Stdout.Write(marshalled)
}
