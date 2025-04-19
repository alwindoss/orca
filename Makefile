
.Phony: build
build:
	go build -o bin/

.Phony: clean
clean:
	rm -rf ./bin
	go clean