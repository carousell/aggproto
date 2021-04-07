
install:
	go install tools/savedescriptor/protoc-gen-savedescriptor.go
	go install tools/registerproto/protoc-gen-registerproto.go
	go install aggproto.go

test_proto:
	protoc --savedescriptor_out=test/testdata/ -I=test/resources/protos test/resources/protos/pkga/a.proto
	protoc --savedescriptor_out=test/testdata/ -I=test/resources/protos test/resources/protos/pkgb/b.proto

test: test_proto
	go test ./...

clean:
	rm -rf examples/registry
examples: clean install
	protoc --registerproto_out=registry_path=examples/registry:. -I=examples/protos examples/protos/listing/listing.proto
	protoc --registerproto_out=registry_path=examples/registry:. -I=examples/protos examples/protos/listing_comments/comments.proto
	aggproto sync --api_specs_path examples/specs --registry_path examples/registry --go_out_path examples/goOut --proto_out_path examples/protoOut
	protoc --go_out=examples/goOut --go-grpc_out=examples/goOut --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --proto_path=examples/protoOut simple_mask_v1/masked_listing.proto
	protoc --go_out=examples/goOut --go-grpc_out=examples/goOut --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --proto_path=examples/protos listing/listing.proto
	protoc --go_out=examples/goOut --go-grpc_out=examples/goOut --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --proto_path=examples/protos listing_comments/comments.proto
