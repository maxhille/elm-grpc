proto = build/example.pb.go

# all: $(js)
all: $(proto)

watch:
	while true; do \
		make all; \
		inotifywait -qre close_write .; \
	done

clean:
	rm -f build/*

$(proto): example.proto
	mkdir -p build
	protoc \
	--go_out=plugins=grpc:build \
  example.proto \
