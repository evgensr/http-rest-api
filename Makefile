.PHONY: build
build:
	go build -v ./cmd/apiserver


.PHONY: test	
test:
	go test -v -race -timeout 30s ./...

.PHONY: pg	
pg:
	docker-compose -f docker-compose-pg-only.yml up --build

.DEFAULT_GOAL := build


