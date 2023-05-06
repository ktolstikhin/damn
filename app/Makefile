.PHONY: build
build:
	@go build -o=./bin/damn ./cmd/damn

.PHONY: run
run: build
	@./bin/damn

.PHONY: test
test:
	@go test -v ./... -count=1
