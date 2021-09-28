DKR_DIR := $(PROJECT_ROOT)/build/docker
DKR_FILE := $(DKR_DIR)/Dockerfile

DKR_IMG_TAG := $(APP_NAME):$(GIT_COMMIT)

.PHONY: docker.build
docker.build: 
	@docker build -f $(DKR_FILE) --build-arg ARCH=$(DKR_ARCH) -t $(DKR_IMG_TAG) $(PROJECT_ROOT)
