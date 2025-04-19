
.Phony: build
build:
	go build -o bin/

.Phony: clean
clean:
	rm -rf ./bin
	rm -rf ./dist
	go clean

.Phony: release
release: clean
	go tool goreleaser release