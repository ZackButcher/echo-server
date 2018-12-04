VENDOR := third_party/protobuf/

PROTO_SOURCES = $(wildcard *.proto)
PROTO_GO_SOURCES = $(patsubst %.proto,%.pb.go,$(PROTO_SOURCES))
DESCRIPTOR := echo.proto-descriptor

.PHONY: all
all: $(DESCRIPTOR)
	@echo "All protos done"

${DESCRIPTOR}:
	protoc \
		-I ./ \
		-I $(VENDOR) \
		--go_out=plugins=grpc:. \
		--descriptor_set_out=$(DESCRIPTOR) --include_imports \
		$(PROTO_SOURCES)

clean:
	rm -f $(PROTO_GO_SOURCES) $(DESCRIPTOR)
	@echo "Cleaned ${PROTO_GO_SOURCES} $(DESCRIPTOR)"