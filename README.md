# timer_apiserver
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
    - [Learn More about Usage of GNU make](#learn-more-about-usage-of-gnu-make)
  - [API Doc](#api-doc)
    - [RESTful API](#restful-api)
    - [gRPC API](#grpc-api)

## Get Started

### Prerequisites
- Go1.17+
- Docker
- GNU make

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

### Learn More about Usage of GNU make
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
The service definition file is like below.

```proto3
syntax = "proto3";

package proto;

service Resource {
    rpc ListPendingTimers(ListPendingTimersReq) returns (ListPendingTimersResp) {}
}

message TimerInfo {
    string name = 1;
    string trigger_at = 2;
}

message ListPendingTimersReq {
}

message ListPendingTimersResp {
    repeated TimerInfo items = 1;
}
```
A timer will be considered as "pending" if it is created, not deleted, and not triggerred yet.
