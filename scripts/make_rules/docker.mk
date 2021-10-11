DKR_DIR := $(PROJECT_ROOT)/build/docker
DKR_FILE := $(DKR_DIR)/Dockerfile
DKR_COMPOSE_FILE := $(DKR_DIR)/docker-compose.yml

DKR_IMG_TAG := $(APP_NAME):$(GIT_COMMIT)

## docker.build: Build the docker image
.PHONY: docker.build
docker.build: 
	@docker build -f $(DKR_FILE) --build-arg ARCH=$(DKR_ARCH) -t $(DKR_IMG_TAG) $(PROJECT_ROOT)

## docker.compose.up: Bring up Timer API Server and database with docker compose
.PHONY: docker.compose.up
docker.compose.up: 
	@cp $(PROJECT_ROOT)/config/example.yml $(PROJECT_ROOT)/config/config.yml
	@GIT_COMMIT=$(GIT_COMMIT) docker compose -f $(DKR_COMPOSE_FILE) -p $(APP_NAME) up --build --detach --force-recreate
	@$(MAKE) mysql.migrate.up

## docker.compose.down: Bring down containers brought up by phony docker.compose.up
.PHONY: docker.compose.down
docker.compose.down: 
	@GIT_COMMIT=$(GIT_COMMIT) docker compose -f $(DKR_COMPOSE_FILE) -p $(APP_NAME) down

## docker.compose.logs: Show logs inside containers brought up by phony docker.compose.up
.PHONY: docker.compose.logs
docker.compose.logs: 
	@GIT_COMMIT=$(GIT_COMMIT) docker compose -f $(DKR_COMPOSE_FILE) -p $(APP_NAME) logs -f

