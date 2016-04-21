PROTO_FILES=$(wildcard *.proto)
PB_FILES=$(patsubst %.proto,%.pb.go,$(PROTO_FILES))

all: $(PB_FILES)

$(PB_FILES): %.pb.go: %.proto
	@cd ../.. && protoc --go_out=. $(patsubst %,zvelo.io/msg/%,$(PROTO_FILES))

.PHONY: all
