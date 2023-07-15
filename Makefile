proto:
	protoc -I pb --go_out=pb --go-grpc_out=pb pb/service.proto
run:
	go run cmd/main.go
