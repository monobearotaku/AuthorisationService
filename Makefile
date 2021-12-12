gen:
	protoc --go_out=proto/compiled --go-grpc_out=proto/compiled  proto/raw/authorization/authorization.proto

runAS:
	go run service/AuthorizationServer/main.go

runC:
	go run client/main.go