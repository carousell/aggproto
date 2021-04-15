
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
	# generate pb.go and grpc.pb.go files for dependent protos used
	protoc --go_out=examples/goOut --go-grpc_out=examples/goOut --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --proto_path=examples/protos listing/listing.proto
	protoc --go_out=examples/goOut --go-grpc_out=examples/goOut --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --proto_path=examples/protos listing_comments/comments.proto
	protoc --go_out=examples/goOut --go-grpc_out=examples/goOut --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --proto_path=examples/protos media/media.proto
	protoc --go_out=examples/goOut --go-grpc_out=examples/goOut --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --proto_path=examples/protos wallet/wallet.proto
	# register internal apis to registry
	protoc --registerproto_out=registry_path=examples/registry:. -I=examples/protos examples/protos/listing/listing.proto
	protoc --registerproto_out=registry_path=examples/registry:. -I=examples/protos examples/protos/listing_comments/comments.proto
	protoc --registerproto_out=registry_path=examples/registry:. -I=examples/protos examples/protos/media/media.proto
	protoc --registerproto_out=registry_path=examples/registry:. -I=examples/protos examples/protos/wallet/wallet.proto
	# sync specs to generate protos and corresponding generated code
	aggproto sync --api_specs_path examples/specs --registry_path examples/registry --go_out_path examples/goOut --proto_out_path examples/protoOut
	# generate pb.go and grpc.pb.go files for generated protos - server code
	protoc --go_out=examples/goOut --go-grpc_out=examples/goOut --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --proto_path=examples/protoOut static_primitives_v1/mock_listing.proto
	protoc --go_out=examples/goOut --go-grpc_out=examples/goOut --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --proto_path=examples/protoOut simple_mask_v1/masked_listing.proto
	protoc --go_out=examples/goOut --go-grpc_out=examples/goOut --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --proto_path=examples/protoOut union_mask_v1/masked_listing_w_wallet.proto
	protoc --go_out=examples/goOut --go-grpc_out=examples/goOut --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --proto_path=examples/protoOut alias_union_mask_v1/masked_listing_w_wallet.proto
	protoc --go_out=examples/goOut --go-grpc_out=examples/goOut --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --proto_path=examples/protoOut inferred_input_v1/masked_listing_w_media.proto
