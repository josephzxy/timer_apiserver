MYSQL_USER ?= root
MYSQL_PWD ?= root
MYSQL_HOST ?= localhost
MYSQL_PORT ?= 3306
MYSQL_DB ?= test

MYSQL_DSN = "mysql://$(MYSQL_USER):$(MYSQL_PWD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DB)"
MYSQL_MIGRATION_DIR = $(PROJECT_ROOT)/database/migrations
MYSQL_MIGRATION_SRC = "file://$(MYSQL_MIGRATION_DIR)"

MYSQL_CONTAINER_NAME := "timer_mariadb"

MYSQL_MK_PREFIX := "MySQL:"

## mysql.migrate.mkdir: Ensure the migrations directory exists
.PHONY: mysql.migrate.mkdir
mysql.migrate.mkdir:
	@echo "=======> $(MYSQL_MK_PREFIX) ensuring migrations directory exists"
	@mkdir -p $(MYSQL_MIGRATION_DIR)

# Below provides handy wrapping for golang-migrate cli
# For detailed usage of golang-migrate cli, see https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#usage


# create [-ext E] [-dir D] [-seq] [-digits N] [-format] NAME
#                Create a set of timestamped up/down migrations titled NAME, in directory D with extension E.
#                Use -seq option to generate sequential up/down migrations with N digits.
#                Use -format option to specify a Go time format string.

## mysql.migrate.create.%: Create MySQL migration with the given name
.PHONY: mysql.migrate.create.%
mysql.migrate.create.%: mysql.migrate.mkdir tools.verify.go-migrate
	@echo "=======> $(MYSQL_MK_PREFIX) creating new MySQL migration: $*"
	@migrate -database $(MYSQL_DSN) create -ext .sql -dir $(MYSQL_MIGRATION_DIR) $*

# goto V       Migrate to version V
## mysql.migrate.goto.%: Update MySQL migration to the given version
.PHONY: mysql.migrate.goto.%
mysql.migrate.goto.%: mysql.migrate.mkdir tools.verify.go-migrate
	@echo "=======> $(MYSQL_MK_PREFIX) migrating to version $*"
	@migrate -database $(MYSQL_DSN) -source $(MYSQL_MIGRATION_SRC) goto $*

# up [N]       Apply all or N up migrations
## mysql.migrate.up: Apply all up migrations
.PHONY: mysql.migrate.up
mysql.migrate.up: mysql.migrate.mkdir tools.verify.go-migrate
	@echo "=======> $(MYSQL_MK_PREFIX) applying all up migrations"
	@migrate -database $(MYSQL_DSN) -source $(MYSQL_MIGRATION_SRC) up

## mysql.migrate.up.%: Apply N up migrations. N is the given integer
.PHONY: mysql.migrate.up.%
mysql.migrate.up.%: mysql.migrate.mkdir tools.verify.go-migrate
	@echo "=======> $(MYSQL_MK_PREFIX) applying $* up migrations"
	@migrate -database $(MYSQL_DSN) -source $(MYSQL_MIGRATION_SRC) up $*

# down [N]     Apply all or N down migrations
## mysql.migrate.down: Apply all down migrations
.PHONY: mysql.migrate.down
mysql.migrate.down: mysql.migrate.mkdir tools.verify.go-migrate
	@echo "=======> $(MYSQL_MK_PREFIX) applying all down migrations"
	@migrate -database $(MYSQL_DSN) -source $(MYSQL_MIGRATION_SRC) down

## mysql.migrate.down.%: Apply N down migrations. N is the given integer
.PHONY: mysql.migrate.down.%
mysql.migrate.down.%: mysql.migrate.mkdir tools.verify.go-migrate
	@echo "=======> $(MYSQL_MK_PREFIX) applying $* down migrations"
	@migrate -database $(MYSQL_DSN) -source $(MYSQL_MIGRATION_SRC) down $*
	
# drop         Drop everything inside database
## mysql.migrate.drop: Drop everything inside the database
.PHONY: mysql.migrate.drop
mysql.migrate.drop: mysql.migrate.mkdir tools.verify.go-migrate
	@echo "=======> $(MYSQL_MK_PREFIX) [!!DANGREROUS!!] dropping entire database schema"
	@migrate -database $(MYSQL_DSN) -source $(MYSQL_MIGRATION_SRC) drop 

# force V      Set version V but don't run migration (ignores dirty state)
## mysql.migrate.force.%: Update the MySQL migration version to the given version by force without actually running migrations
.PHONY: mysql.migrate.force.%
mysql.migrate.force.%: mysql.migrate.mkdir tools.verify.go-migrate
	@echo "=======> $(MYSQL_MK_PREFIX) forcing migration version to $* without actually running migrations"
	@migrate -database $(MYSQL_DSN) -source $(MYSQL_MIGRATION_SRC) force $*

# version      Print current migration version
## mysql.migrate.version: Print the current version of MySQL migration
.PHONY: mysql.migrate.version
mysql.migrate.version: mysql.migrate.mkdir tools.verify.go-migrate
	@echo "=======> $(MYSQL_MK_PREFIX) printting current migration version"
	@migrate -database $(MYSQL_DSN) -source $(MYSQL_MIGRATION_SRC) version

## mysql.docker.start: Starts a MySQL docker container 
.PHONY: mysql.docker.start
mysql.docker.start:
	@echo "=======> $(MYSQL_MK_PREFIX) starting mysql container"
	@docker run -d --rm --name $(MYSQL_CONTAINER_NAME) -p $(MYSQL_PORT):3306 \
	 -e MYSQL_DATABASE=$(MYSQL_DB) \
	 -e MARIADB_ROOT_PASSWORD=$(MYSQL_PWD) \
	 -e MYSQL_PWD=$(MYSQL_PWD) \
	 mariadb
	@echo "=======> $(MYSQL_MK_PREFIX) waiting 5 seconds for the container to fully start"
	@sleep 5
	@$(MAKE) mysql.migrate.up

## mysql.docker.stop: Stops (and removes) the MySQL docker container launched earlier
.PHONY: mysql.docker.stop
mysql.docker.stop:
	@echo "=======> $(MYSQL_MK_PREFIX) stopping mysql container"
	@docker stop $(MYSQL_CONTAINER_NAME)

