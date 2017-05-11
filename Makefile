SWAGGER_UI_VERSION := 3.0.9
FIRST_GOPATH       := $(firstword $(subst :, ,$(GOPATH)))
PROTO_FILES        := $(wildcard *.proto)
GO_PB_FILES        := $(patsubst %.proto,%.pb.go,$(PROTO_FILES))
PY_PB_FILES        := $(patsubst %.proto,%_pb2.py,$(PROTO_FILES))
GRPC_GATEWAY_PROTO_FILES := query_api.proto
GRPC_GATEWAY_FILES := $(patsubst %.proto,%.pb.gw.go,$(GRPC_GATEWAY_PROTO_FILES))

default: go grpc-gateway swagger.json static.go
go: $(GO_PB_FILES)
python: $(PY_PB_FILES)
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
-I$(FIRST_GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
$(1)
endef

define protoc-go
--gozvelo_out=plugins=grpc:. \
$(patsubst %,zvelo/msg/%,$(PROTO_FILES))
endef

define protoc-grpc-gateway
--grpc-gateway_out=logtostderr=true:. \
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

swagger.json: $(GRPC_GATEWAY_PROTO_FILES)
	$(call wrap-cmd,$(call wrap-protoc,$(protoc-swagger)))
	@mv $(patsubst %.proto,%.swagger.json,$<) swagger.json

.swagger-ui.tar.gz:
	curl -sLo $@ https://github.com/swagger-api/swagger-ui/archive/v$(SWAGGER_UI_VERSION).tar.gz

.swagger-ui: .swagger-ui.tar.gz
	@rm -rf $@
	tar -zxf $<
	mv swagger-ui-$(SWAGGER_UI_VERSION) $@
	sed -i 's|^url: ".*"$$|url: "/v1/swagger.json"|' $@/swagger-config.yaml
	cd $@ && npm install && npm run build
	@touch $@

static.go: swagger.json .swagger-ui
	mkdir -p static/v1
	cp -a swagger.json static/v1
	cp -a .swagger-ui/dist static/swagger-ui
	esc -o static.go -pkg msg -prefix static static
	rm -rf static

$(PY_PB_FILES): %_pb2.py: %.proto
	$(call wrap-cmd,$(protoc-python))

clean:
	rm -rf $(GO_PB_FILES) $(PY_PB_FILES) $(GRPC_GATEWAY_FILES) swagger.json .swagger-ui.tar.gz .swagger-ui static.go

.PHONY: default clean go python grpc-gateway
