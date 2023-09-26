lint:
	golangci-lint run -c .golangci.yml
generate-server:
	go generate back/api/gen.go
up-docker-compose:
	docker compose -f ./docker-compose.yaml -p back up --remove-orphans --force-recreate -d --build
down-docker-compose:
	docker compose down --remove-orphans