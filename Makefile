proto = build/example.pb.go build/Example.elm build/example_pb.js \
 	build/example_pb_service.js
elm = build/elm.js
js = build/index.js build/index.js.map

all: $(proto) $(elm) $(js)

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

$(js): $(elm) $(proto) glue.js webpack.config.js
	mkdir -p build
	webpack

$(elm): Main.elm build/Example.elm
	mkdir -p build
	elm make Main.elm --output build/elm.js
