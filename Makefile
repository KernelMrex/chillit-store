build:
	go build -v ./cmd/chillitstore/.

run:
	go run -v ./cmd/chillitstore/. -config_path=./configs/config.yaml

run_dev:
	go run -v ./cmd/chillitstore/. -config_path=./configs/config.yaml.devel

test:
	go test -v -race ./...

.DEFAULT_GOAL := run