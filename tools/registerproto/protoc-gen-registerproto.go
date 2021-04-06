package main

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/carousell/aggproto/pkg/registry/parser"
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
	var p parser.Parser
	parameters := req.GetParameter()
	for _, group := range strings.Split(parameters, ",") {
		kv := strings.Split(group, "=")
		if len(kv) > 1 {
			if kv[0] == "registry_path" {
				p = parser.LoadParser(kv[1])
			}
		}
	}
	if p == nil {
		panic("please pass registry_path")
	}
	for _, fdp := range req.ProtoFile {
		err := p.AddProtoContainer(fdp)
		if err != nil {
			panic(err)
		}
	}
	err = p.SaveRegistry()
	if err != nil {
		panic(err)
	}
	marshalled, err := proto.Marshal(resp)
	if err != nil {
		panic(err)
	}
	_, _ = os.Stdout.Write(marshalled)

}
