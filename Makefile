PROTO_FILES=$(sort $(wildcard *.proto))

all: msg

msg: .msg

.msg: $(PROTO_FILES)
	protoc -I ../.. --go_out=plugins=grpc:go-msg $(patsubst %,../../zvelo.io/msg/%,$(PROTO_FILES))
	@touch .msg

.PHONY: all msg
