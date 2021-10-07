# ⏳ Timer API Server
[![codecov](https://codecov.io/gh/josephzxy/timer_apiserver/branch/develop/graph/badge.svg?token=AU3193CPC3)](https://codecov.io/gh/josephzxy/timer_apiserver)

`timer_apiserver` is a demo Golang project that implements both external facing RESTful APIs and internal facing gRPC APIs for managing resource `timer`.
Resource `timer` will stored in `MySQL`

**Major Tech stacks:**
Golang, Gin, gRPC, MySQL, GORM

- [timer_apiserver](#timer_apiserver)
  - [Get Started](#get-started)
    - [Prerequisites](#prerequisites)
    - [Build Local Executable](#build-local-executable)
    - [Build Docker Image](#build-docker-image)
    - [Deploy Locally with Docker Compose for Development](#deploy-locally-with-docker-compose-for-development)
    - [MySQL Database Schema Migration](#mysql-database-schema-migration)
      - [Quick Example](#quick-example)
      - [Use Custom MySQL Config](#use-custom-mysql-config)
    - [Show help message for make](#show-help-message-for-make)
  - [API Doc](#api-doc)
    - [RESTful API](#restful-api)
    - [gRPC API](#grpc-api)

## Get Started

### Prerequisites
- [Go1.17+](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started)
- [GNU make](https://www.gnu.org/software/make/)
- [protoc](https://grpc.io/docs/protoc-installation/#install-using-a-package-manager)

### Build Local Executable
For the host OS and ARCH.
```bash
# get into the project root directory
make
```
For a certain platform, specify `PLATFORM`.
```bash
make PLATFORM=linux_amd64
```

### Build Docker Image
The docker image OS is default to `linux`. 

For host ARCH
```bash
make docker
```
For other ARCHs, specify `DKR_ARCH`.
```bash
make docker DKR_ARCH=amd64
```

### Deploy Locally with Docker Compose for Development
```bash
make docker.compose.up
make docker.compose.down
```

### MySQL Database Schema Migration
Some handy wrappings are provided in Make Phonies with prefix `mysql.migrate.*`.

#### Quick Example
Below is a quick example for creating and applying a migration. For more details, see [mysql.mk](./scripts/make_rules/mysql.mk)
```bash
# create migration for adding table foo
$ make mysql.migrate.create.add_table_foo
=======> MySQL: ensuring migrations directory exists
=======> MySQL: creating new MySQL migration: foo
/some/dir/timer_apiserver/database/migrations/20210930054712_foo.up.sql
/some/dir/timer_apiserver/database/migrations/20210930054712_foo.down.sql

# fill in desired SQL statements into .sql files newly generated above

# apply up migrations
$ make mysql.migrate.up
=======> MySQL: applying all up migrations
20210930054712/u add_table_foo (33.492584ms)

# now table foo is created

# to undo it, apply down migrations
$ make mysql.migrate.down
=======> MySQL: applying all down migrations
Are you sure you want to apply all down migrations? [y/N]
y
Applying all down migrations
20210930054712/d add_table_foo (21.437875ms)

# now table foo is gone
```
#### Use Custom MySQL Config
By default, dummy configs listed in [mysql.mk](./scripts/make_rules/mysql.mk) are used to connect to MySQL.
```make
MYSQL_USER ?= root
MYSQL_PWD ?= root
MYSQL_HOST ?= localhost
MYSQL_PORT ?= 3306
MYSQL_DB ?= test
```
To use custom config, specify environment variables when invoking Make
```bash
MYSQL_PORT=3307 make mysql.migrate.up
```

### Show help message for make
```bash
make help
```

## API Doc

### RESTful API

View swagger doc as a webpage by
```bash
make swagger.serve
```
Or check out [swagger.yml](api/rest/swagger/swagger.yml) directly.

### gRPC API
Please check out [timer.proto](api/grpc/timer.proto).
A timer will be considered as "pending" if it is created, not deleted, and not triggerred yet.
