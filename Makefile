all: go ruby

go:
	protoc -I=. ./agent_api.proto --go_out=plugins=grpc:.


ruby:
	grpc_tools_ruby_protoc -I ./ --ruby_out=. --grpc_out=. ./agent_api.proto
