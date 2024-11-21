.PHONY: postgres migrate
postgres:
	docker run --rm -ti --network host -e POSTGRES_PASSWORD=example postgres

DB_URL="postgres://postgres:example@192.168.106.2/postgres?sslmode=disable"
migrate:
	migrate -source file://./migrations -database $(DB_URL) up

migrate-down:
	migrate -source file://./migrations -database $(DB_URL) down
