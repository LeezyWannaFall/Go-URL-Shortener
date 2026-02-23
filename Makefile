.PHONY: run test down

run:
	sudo docker compose up -d --build

down:
	sudo docker compose down

test:
	go test -v ./...