# demo.mk provides handy phonies for trying out timer_apiserver

DEMO_HOST = localhost
DEMO_CONTAINER_HOST = host.docker.internal

DEMO_REST_PORT = 8081
DEMO_REST_CONTAINER_ADDR = http://$(DEMO_CONTAINER_HOST):$(DEMO_REST_PORT)

DEMO_GRPC_PORT = 8082
DEMO_GRPC_CONTAINER_ADDR = $(DEMO_CONTAINER_HOST):$(DEMO_GRPC_PORT)

# demo clients
HTTPIE := docker run -it --rm --add-host host.docker.internal:host-gateway alpine/httpie
MYSQL := docker run -it --rm --network=host mariadb mysql
GRPC_CLI := docker run -it --rm --add-host host.docker.internal:host-gateway namely/grpc-cli

DEMO_REST_PUT_NAME ?= "new name"
DEMO_REST_PUT_TRIGGER_AT ?= "2030-01-01T00:00:00+08:00"

# demo for grpc server
.PHONY: demo.grpc.getallpending
demo.grpc.getallpending:
	$(GRPC_CLI) call $(DEMO_GRPC_CONTAINER_ADDR) timer.Timer.GetAllPendingTimers ""

# demo for rest server
.PHONY: demo.rest.get.all
demo.rest.get.all:
	$(HTTPIE) GET $(DEMO_REST_CONTAINER_ADDR)/v1/timers

.PHONY: demo.rest.get.%
demo.rest.get.%:
	$(HTTPIE) GET $(DEMO_REST_CONTAINER_ADDR)/v1/timers/$*

.PHONY: demo.rest.post.%
demo.rest.post.%:
	$(HTTPIE) POST $(DEMO_REST_CONTAINER_ADDR)/v1/timers name=$* triggerAt="2022-10-30T07:59:10+08:00"

.PHONY: demo.rest.put.%
demo.rest.put.%:
	$(HTTPIE) PUT $(DEMO_REST_CONTAINER_ADDR)/v1/timers/$* name=$(DEMO_REST_PUT_NAME) triggerAt=$(DEMO_REST_PUT_TRIGGER_AT)

.PHONY: demo.rest.delete.%
demo.rest.delete.%:
	$(HTTPIE) DELETE $(DEMO_REST_CONTAINER_ADDR)/v1/timers/$*

# demo for db
.PHONY: demo.db.show
demo.db.show:
	$(MYSQL) -P 3306 --protocol=tcp -uroot -proot -e 'use test; select * from timer;'
