proto:
	protoc -I pb --go_out=pb --go-grpc_out=pb pb/service.proto
run:
	go run cmd/main.go
swag:
	swag init -g pkg/handlers/adminDashboard.go -o cmd/docs