all: up

up:
	docker compose --env-file ./.env -f .\CI\docker-compose.yml up --build
down:
	docker compose --env-file ./.env -f .\CI\docker-compose.yml down
