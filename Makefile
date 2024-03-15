.PHONY: test
test:
	go test -race -v ./...

.PHONY: benchmark
benchmark:
	go test $(shell pwd)/pkg/dataformats/ -bench=. -benchmem -cpuprofile=cpu.pprof -memprofile=mem.pprof

.PHONY: style-fix
style-fix:
	gofmt -w .

.PHONY: lint
lint:
	golangci-lint run

.PHONY: upgrade
upgrade:
	go mod tidy
	go get -u all ./...