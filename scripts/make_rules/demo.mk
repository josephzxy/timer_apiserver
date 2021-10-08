# demo.mk provides handy phonies for trying out timer_apiserver

DEMO_HOST = localhost
DEMO_CONTAINER_HOST = host.docker.internal

DEMO_REST_PORT = 8081
DEMO_REST_CONTAINER_ADDR = http://$(DEMO_CONTAINER_HOST):$(DEMO_REST_PORT)

DEMO_REST_DEFAULT_TRIGGER_AT = "2022-10-30T07:59:10+08:00"
DEMO_REST_PUT_NAME ?= "new name"
DEMO_REST_PUT_TRIGGER_AT ?= $(DEMO_REST_DEFAULT_TRIGGER_AT)

DEMO_GRPC_PORT = 8082
DEMO_GRPC_CONTAINER_ADDR = $(DEMO_CONTAINER_HOST):$(DEMO_GRPC_PORT)

# demo clients
HTTPIE := docker run -it --rm --add-host host.docker.internal:host-gateway alpine/httpie
MYSQL := docker run -it --rm --network=host mariadb mysql
GRPC_CLI := docker run -it --rm --add-host host.docker.internal:host-gateway namely/grpc-cli

DEMO_MK_PREFIX := "Demo:"

# demo for grpc server
.PHONY: demo.grpc.getallpending
demo.grpc.getallpending:
	@echo "=======> $(DEMO_MK_PREFIX) [GRPC] getting all pending timers"
	$(GRPC_CLI) call $(DEMO_GRPC_CONTAINER_ADDR) timer.Timer.GetAllPendingTimers ""
	@echo

# demo for rest server
.PHONY: demo.rest.get.all
demo.rest.get.all:
	@echo "=======> $(DEMO_MK_PREFIX) [REST] getting all timers"
	$(HTTPIE) GET $(DEMO_REST_CONTAINER_ADDR)/v1/timers
	@echo

.PHONY: demo.rest.get.%
demo.rest.get.%:
	@echo "=======> $(DEMO_MK_PREFIX) [REST] getting timer $*"
	$(HTTPIE) GET $(DEMO_REST_CONTAINER_ADDR)/v1/timers/$*
	@echo

.PHONY: demo.rest.post.%
demo.rest.post.%:
	@echo "=======> $(DEMO_MK_PREFIX) [REST] creating timer $*"
	$(HTTPIE) POST $(DEMO_REST_CONTAINER_ADDR)/v1/timers name=$* triggerAt=$(DEMO_REST_DEFAULT_TRIGGER_AT)
	@echo

.PHONY: demo.rest.put.%
demo.rest.put.%:
	@echo "=======> $(DEMO_MK_PREFIX) [REST] updating timer $*"
	$(HTTPIE) PUT $(DEMO_REST_CONTAINER_ADDR)/v1/timers/$* name=$(DEMO_REST_PUT_NAME) triggerAt=$(DEMO_REST_PUT_TRIGGER_AT)
	@echo

.PHONY: demo.rest.delete.%
demo.rest.delete.%:
	@echo "=======> $(DEMO_MK_PREFIX) [REST] deleting timer $*"
	$(HTTPIE) DELETE $(DEMO_REST_CONTAINER_ADDR)/v1/timers/$*
	@echo

# demo for db
.PHONY: demo.db.show
demo.db.show:
	@echo "=======> $(DEMO_MK_PREFIX) [DB] displaying current content of table"
	$(MYSQL) -P 3306 --protocol=tcp -uroot -proot -e 'use test; select * from timer;'
	@echo
