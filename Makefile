proto = build/example.pb.go build/Example.elm
elm = build/index.html

all: $(proto) $(elm)

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
	--elm_out=plugins=grpc:build \
  example.proto \

$(elm): Main.elm build/Example.elm
	elm make Main.elm --output build/index.html
