LOCAL_BIN := $(CURDIR)/bin
PROTO_PATH := $(CURDIR)/api/user/v1
PKG_PROTO_PATH := $(CURDIR)/proto

# определяем, Windows или нет
ifeq ($(OS),Windows_NT)
  EXE := .exe
else
  EXE :=
endif

.bin-deps:
	@echo Installing plugins...
	@GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.protoc-generate:
	mkdir -p $(PKG_PROTO_PATH)
	protoc \
		--proto_path=$(PROTO_PATH) \
		--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go$(EXE) \
		--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc$(EXE) \
		--go_out=$(PKG_PROTO_PATH) --go_opt=paths=source_relative \
		--go-grpc_out=$(PKG_PROTO_PATH) --go-grpc_opt=paths=source_relative \
		$(PROTO_PATH)/service.proto \
		$(PROTO_PATH)/messages.proto

.tidy:
	GOBIN=$(LOCAL_BIN) go mod tidy

generate: .bin-deps .protoc-generate .tidy

build:
	go build -o $(LOCAL_BIN) ./cmd/main.go

.PHONY: .bin-deps .protoc-generate .tidy generate build
