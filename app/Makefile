.PHONY: build
build:
	go build -o=./bin/api ./cmd/api
	go build -o=./bin/damn ./cmd/damn

.PHONY: run
run: build
	./bin/api -addr=:8000 -debug

.PHONY: test
test:
	go test -v ./... -count=1
