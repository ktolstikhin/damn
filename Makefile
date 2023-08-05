.PHONY: build
build:
	docker compose build

build/web:
	docker compose build web

build/api:
	docker compose build api

.PHONY: run
run:
	docker compose up

.PHONY: deploy
deploy:
	docker compose \
		-f docker-compose.yml \
		-f docker-compose.production.yml \
		up -d

.PHONY: stop
stop:
	docker compose down
