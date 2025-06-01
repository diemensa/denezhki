.PHONY: test build up all

all: test up

test:
	go test ./...

up:
	docker-compose up --build
