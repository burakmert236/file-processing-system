OUT_DIR=../generated
PROTO_FILES=$(wildcard *.proto)

all:
	mkdir -p $(OUT_DIR)
	protoc --go_out=$(OUT_DIR) \
		--go_opt=paths=source_relative \
		$(PROTO_FILES)