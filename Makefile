start:
	docker compose up -d

start-with-build:
	docker compose up --build -d

stop:
	docker compose down

logs:
	docker compose logs -f pack-service

backend-run:
	go build ./cmd/app && ./app -c config.toml

test:
	go clean -testcache && go test -race -cover ./...

setup-tools:
	go install github.com/vektra/mockery/v2@v2.53.3

mockery:
	mockery --all --case=underscore --keeptree

docker-test:
	docker-compose -f docker-compose.test.yml run --rm --build test; \
	$(MAKE) stop

lint:
	golangci-lint run