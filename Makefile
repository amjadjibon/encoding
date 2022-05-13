VERSION = "0.0.1"
change-version:
	@echo $(VERSION)>VERSION
	@echo "package constant\n\n// Version constant of encoding\nconst Version = \"$(VERSION)\"">constant/version.go
	@git add VERSION
	@git add constant/version.go
	@git commit -m "v$(VERSION)"
	@git tag -a "v$(VERSION)" -m "v$(VERSION)"
	@git push origin
	@git push origin "v$(VERSION)"

update-module:
	go get github.com/mkawserm/abesh
	go get google.golang.org/protobuf/proto
	go get github.com/vmihailenco/msgpack/v5
	go get github.com/pelletier/go-toml/v2
	go get github.com/clbanning/mxj/v2

run-example:
	go run example/main/main.go embedded run

test:
	go test ./... -v -cover -coverprofile=coverage.out -covermode=atomic && go tool cover -html=coverage.out

update-protoc:
	@protoc \
		-I=./proto \
		--go_opt=module=github.com/amjadjibon/encoding \
		--go_out=. \
		./proto/test.proto
