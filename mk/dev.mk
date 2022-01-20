## Build and start the service in development mode (detached)
run: build-dev "docker-compose-up -d"

## Build and start the service in development mode (attached)
start: build-dev docker-compose-up

## Stop running services
stop: docker-compose-down

.PHONY: run start stop

## Build develoment docker image
build-dev: docker-compose-build

## Run docker compose commands with the project configuration
docker-compose-%:
	docker-compose -f docker-compose.yml \
		--project-directory . \
		$*