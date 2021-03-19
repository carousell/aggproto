
install:
	go install tools/protoc-gen-savedescriptor.go

test_proto:
	protoc --savedescriptor_out=test/testdata/ -I=test/resources/protos test/resources/protos/pkga/a.proto
	protoc --savedescriptor_out=test/testdata/ -I=test/resources/protos test/resources/protos/pkgb/b.proto

test: test_proto
	go test ./...