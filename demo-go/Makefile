.PHONY: run
run:
	go run ./cmd/demo/main.go --config ../config.yaml

.PHONY: run-with-reload
run-with-reload:
	`go env GOPATH`/bin/air --build.cmd "go build -o ./bin/demo ./cmd/demo/main.go" --build.bin "./bin/demo --config ../config.yaml"

.PHONY: run-load
run-load:
	./load-generator.sh
