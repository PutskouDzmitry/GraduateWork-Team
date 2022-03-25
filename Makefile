.PHONY: build_docker_postgres
build_docker_postgres:
	docker build -t docker_db_postgres -f ./docker/db.dockerfile .

.PHONY: build_db
build_db: build_docker_postgres

.PHONY: run_docker_postgres
run_docker_postgres:
	docker run -p 5432:5432 -e POSTGRES_PASSWORD=password -d docker_db_postgres

.PHONY: run_docker_redis
run_docker_redis:
	docker run --name my-redis-db -p 6379:6379 -d redis

.PHONY: run_db
run_db: run_docker_postgres run_docker_redis

.PHONY: build_server
build_server:
	docker build -t docker_server -f ./docker/db.dockerfile .

.PHONY: build_client
build_client:
	npm i

.PHONE: run_client
run_client:
	npm start

.PHONY: build
build: build_db

.PHONY: run
run: run_db