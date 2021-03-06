.PHONY: build
build:
	go build -v ./cmd/apiserver


.PHONY: test	
test:
	go test -v -race -timeout 30s ./...

.PHONY: pg	
pg:
	docker-compose -f docker-compose-pg-only.yml up --build

open-adminer:
	open http://localhost:8081/?pgsql=db&username=postgres&db=restapi_dev&ns=public

run:
	go build -v ./cmd/apiserver && ./apiserver

.DEFAULT_GOAL := build


