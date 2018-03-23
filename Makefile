proto = build/example.pb.go build/Example.elm build/example_service_pb.js build/example_service_pb_service.js
elm = build/index.html
js = app/index.js app/index.js.map

all: $(proto) $(elm)

clean:
	rm -rf build

$(proto): example.proto
	mkdir -p build
	protoc \
	--go_out=plugins=grpc:build \
	--elm_out=plugins=grpc:build \
	--plugin=protoc-gen-js_service=./node_modules/.bin/protoc-gen-js_service \
	--js_out=import_style=commonjs,binary:build \
	--js_service_out=build \
  example.proto \

$(elm): Main.elm build/Example.elm
	mkdir -p build
	elm make Main.elm --output build/index.html
