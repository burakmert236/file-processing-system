OUT_DIR=../generated
PROTO_FILES=$(wildcard *.proto)

all:
	mkdir -p $(OUT_DIR)
	protoc \
		--proto_path=../proto \
		--go_out=$(OUT_DIR) \
		--go_opt=paths=source_relative \
		commands/commands.proto \
		events/events.proto \
		common/common.proto