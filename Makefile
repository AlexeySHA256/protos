PROTO_VERSION ?= v1
PROTO_GEN_DIR ?= ./gen
PROTOS := $(wildcard $(PROTO_VERSION)/*.proto)
GO_PACKAGE=github.com/AlexeySHA256/ssopb
GO_OPTS := $(foreach f,$(PROTOS),--go_opt=M$(f)=$(GO_PACKAGE))
GO_GRPC_OPTS := $(foreach f,$(PROTOS),--go-grpc_opt=M$(f)=$(GO_PACKAGE))

build/go:
	protoc \
	  --go_out=$(PROTO_GEN_DIR) --go_opt=paths=source_relative \
	  --go-grpc_out=$(PROTO_GEN_DIR) --go-grpc_opt=paths=source_relative \
	  $(GO_OPTS) $(GO_GRPC_OPTS) $(PROTOS)
