gen:
	protoc --go_out=proto/compiled --go-grpc_out=proto/compiled  proto/raw/authorization/authorization.proto

run_server:
	go run service/main.go

run_client:
	go run client/main.go