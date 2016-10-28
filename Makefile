PROTO_FILES  := $(wildcard *.proto)
GO_PB_FILES  := $(patsubst %.proto,%.pb.go,$(PROTO_FILES))
PY_PB_FILES  := $(patsubst %.proto,%_pb2.py,$(PROTO_FILES))
FIRST_GOPATH := $(firstword $(subst :, ,$(GOPATH)))

default: go
go: $(GO_PB_FILES)
python: $(PY_PB_FILES)

define wrap-cmd
@rm -f ../../zvelo
@ln -sf zvelo.io ../../zvelo
cd ../.. && $(1)
@rm -f ../../zvelo
endef

define wrap-protoc
protoc \
	-I. \
	-I$(FIRST_GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	$(1)
endef

define protoc-go
--go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. \
$(patsubst %,zvelo/msg/%,$(PROTO_FILES))
endef

define protoc-python
python \
	-m grpc.tools.protoc \
	--python_out=. \
	--grpc_python_out=. \
	-I. \
	-I$(FIRST_GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	$(patsubst %,zvelo/msg/%,$(PROTO_FILES))
endef

$(GO_PB_FILES): %.pb.go: %.proto
	$(call wrap-cmd,$(call wrap-protoc,$(protoc-go)))

$(PY_PB_FILES): %_pb2.py: %.proto
	$(call wrap-cmd,$(protoc-python))

clean:
	rm -rf $(GO_PB_FILES) $(PY_PB_FILES)

.PHONY: default clean go python
