PROTO_FILES=$(wildcard *.proto)
GO_PB_FILES=$(patsubst %.proto,%.pb.go,$(PROTO_FILES))
PY_PB_FILES=$(patsubst %.proto,%_pb2.py,$(PROTO_FILES))

default: go
go: $(GO_PB_FILES)
python: $(PY_PB_FILES)

$(GO_PB_FILES): %.pb.go: %.proto
	rm -f ../../zvelo
	ln -sf zvelo.io ../../zvelo
	cd ../.. && protoc --go_out=. $(patsubst %,zvelo/msg/%,$(PROTO_FILES))
	rm -f ../../zvelo

$(PY_PB_FILES): %_pb2.py: %.proto
	rm -f ../../zvelo
	ln -sf zvelo.io ../../zvelo
	cd ../.. && python -m grpc.tools.protoc --python_out=. --grpc_python_out=. -I. $(patsubst %,zvelo/msg/%,$(PROTO_FILES))
	rm -f ../../zvelo

.PHONY: default go python
