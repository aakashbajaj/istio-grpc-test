proto:
	protoc  -I /Users/techno/work/istio-grpc/internal -I /Users/techno/work/istio-grpc/proto --include_imports --descriptor_set_out=/Users/techno/work/istio-grpc/proto/service.protoset --go_out=paths=source_relative:/Users/techno/work/istio-grpc/internal /Users/techno/work/istio-grpc/proto/service.proto --go-grpc_out=/Users/techno/work/istio-grpc/

clean:
	rm pb/*.go

run:
	go run main.go