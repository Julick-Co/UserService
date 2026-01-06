# UserService
Сервис для юзеров


## Генерация protobuf и gRPC
 
Запусти Makefile

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --proto_path=./api --go_out=. --go-grpc_out=. api/user/v1/messages.proto api/user/v1/service.proto
