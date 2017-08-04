all: go ruby

go:
	protoc -I=. ./agent_api.proto --go_out=plugins=grpc:.


ruby:
	protoc -I=. ./agent_api.proto --ruby_out=./ --plugin=protoc-gen-grpc=`which grpc_ruby_plugin`
