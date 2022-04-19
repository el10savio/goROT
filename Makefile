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

lint-dockerfile:
	@echo "lint rot Dockerfile"
	docker run --rm -i hadolint/hadolint < Dockerfile

rot-build:
	@echo "Building rot Docker Image"
	DOCKER_BUILDKIT=1 docker build -t goROT -f Dockerfile .

rot-run:
	@echo "Running Single rot Docker Container"
	docker run -d goROT

codespell:
	@echo "checking rot spellings"
	codespell
