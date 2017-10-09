FIRST_GOPATH             := $(firstword $(subst :, ,$(GOPATH)))
PROTO_FILES              := $(wildcard *.proto)
GO_PB_FILES              := $(patsubst %.proto,%.pb.go,$(PROTO_FILES))
PY_PB_FILES              := $(patsubst %.proto,%_pb2.py,$(PROTO_FILES))
GRPC_GATEWAY_PROTO_FILES := api.proto
GRPC_GATEWAY_FILES       := $(patsubst %.proto,%.pb.gw.go,$(GRPC_GATEWAY_PROTO_FILES))

.PHONY: default
default: go grpc-gateway swagger.json

.PHONY: go
go: $(GO_PB_FILES) $(GRPC_GATEWAY_FILES)

.PHONY: python
python: $(PY_PB_FILES)

.PHONY: grpc-gateway
grpc-gateway: $(GRPC_GATEWAY_FILES)

define wrap-cmd
@rm -f ../../zvelo
@ln -sf zvelo.io ../../zvelo
cd ../.. && $(1)
@rm -f ../../zvelo
endef

define wrap-protoc
protoc \
-I. \
-Izvelo/msg/include \
-I$(FIRST_GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
$(1)
endef

define protoc-go
--gozvelo_out=plugins=grpc:. \
$(patsubst %,zvelo/msg/%,$(PROTO_FILES))
endef

define protoc-grpc-gateway
--grpc-gateway_out=logtostderr=true,request_context=true:. \
$(patsubst %,zvelo/msg/%,$^)
endef

define protoc-swagger
--swagger_out=logtostderr=true:. \
$(patsubst %,zvelo/msg/%,$<)
endef

define protoc-python
python \
-m grpc.tools.protoc \
--python_out=. \
--grpc_python_out=. \
-I. \
$(patsubst %,zvelo/msg/%,$(PROTO_FILES))
endef

$(GO_PB_FILES): %.pb.go: %.proto
	$(call wrap-cmd,$(call wrap-protoc,$(protoc-go)))

$(GRPC_GATEWAY_FILES): %.pb.gw.go: $(GRPC_GATEWAY_PROTO_FILES)
	$(call wrap-cmd,$(call wrap-protoc,$(protoc-grpc-gateway)))

swagger.json: $(GRPC_GATEWAY_PROTO_FILES) $(PROTO_FILES) internal/swagger-patch/main.go
	$(call wrap-cmd,$(call wrap-protoc,$(protoc-swagger)))
	@mv $(patsubst %.proto,%.swagger.json,$<) swagger.json
	go run ./internal/swagger-patch/main.go

$(PY_PB_FILES): %_pb2.py: %.proto
	$(call wrap-cmd,$(protoc-python))

.PHONY: clean
clean:
	rm -rf $(GO_PB_FILES) $(PY_PB_FILES) $(GRPC_GATEWAY_FILES) swagger.json
