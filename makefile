init:
	sh ./init-sunbmodules.sh

gen-api:
	# Go server + GRPC services + descriptor
	protoc \
		-I 3rdparty/api-common-protos	\
		-I 3rdparty/grpc-gateway	\
		-I proto \
		--include_imports \
		--include_source_info \
		--go-grpc_out=module=azinchenko/auth:pb \
		--go_out=module=azinchenko/auth:pb \
		--descriptor_set_out=deploy/gen/descriptor.pb \
		$(wildcard proto/azinchenko/todo/v1/*.proto) \

#		$(wildcard proto/azinchenko/echo/v1/*.proto)


grpcui:
	grpcui -port 9901 -protoset deploy/gen/descriptor.pb -plaintext localhost:8080

run-local:
	go run cmd/catalog/main.go

