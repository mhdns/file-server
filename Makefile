server:
	go run server/*.go

auth_server:
	go run auth_service/service/*.go

gen_auth_proto:
	protoc --proto_path=auth_service/proto --go_out=auth_service/auth_pb --go_opt=paths=source_relative --go-grpc_out=auth_service/auth_pb --go-grpc_opt=paths=source_relative auth_service/proto/*.proto

gen_file_proto:
	protoc --proto_path=file_service/proto --go_out=file_service/file_pb --go_opt=paths=source_relative --go-grpc_out=file_service/file_pb --go-grpc_opt=paths=source_relative file_service/proto/*.proto

.PHONY: server gen_auth_proto gen_file_proto auth_server