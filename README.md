# UserService
Сервис для юзеров

## Генерация protobuf и gRPC

Запусти команды в терминале:

```bash
# Установка генераторов Go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

```bash
$ protoc -I api/v1/user --go_out=. --go-grpc_out=. api/v1/user/user_messages.proto api/v1/user/user_service.proto
