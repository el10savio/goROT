build:
	@echo "Building rot"
	go build -o bin/rot main.go

fmt:
	@echo "go fmt rot"
	go fmt ./...

vet:
	@echo "go vet rot"
	go vet ./...

test:
	@echo "Testing rot"
	go test -v -race --cover ./...

golanglintci:
	@echo "golanglintci rot"
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.42.1 golangci-lint run --out-format tab --enable-all

semgrep:
	@echo "semgrep rot"
	docker run --rm -v "$(shell pwd):/src" returntocorp/semgrep --config=auto

codespell:
	@echo "checking rot spellings"
	codespell
