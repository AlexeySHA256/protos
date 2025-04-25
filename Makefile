PROTOS := $(wildcard proto/*.proto)
GO_PACKAGE=github.com/AlexeySHA256/ssopb
GO_OPTS := $(foreach f,$(PROTOS),--go_opt=M$(f)=$(GO_PACKAGE))
GO_GRPC_OPTS := $(foreach f,$(PROTOS),--go-grpc_opt=M$(f)=$(GO_PACKAGE))

build/go:
	protoc \
	  --go_out=./ --go_opt=paths=source_relative \
	  --go-grpc_out=./ --go-grpc_opt=paths=source_relative \
	  $(GO_OPTS) $(GO_GRPC_OPTS) $(PROTOS)
