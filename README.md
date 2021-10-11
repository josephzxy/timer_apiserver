# ⏳ Timer API Server
[![codecov](https://codecov.io/gh/josephzxy/timer_apiserver/branch/develop/graph/badge.svg?token=AU3193CPC3)](https://codecov.io/gh/josephzxy/timer_apiserver)
[![CircleCI](https://circleci.com/gh/josephzxy/timer_apiserver.svg?style=svg)](https://circleci.com/gh/josephzxy/timer_apiserver)
[![Go Report Card](https://goreportcard.com/badge/github.com/josephzxy/timer_apiserver)](https://goreportcard.com/report/github.com/josephzxy/timer_apiserver)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/josephzxy/timer_apiserver.svg)](https://github.com/josephzxy/timer_apiserver)

Timer API Server is a demo Golang project that implements RESTful APIs and gRPC APIs for Timer.

**Major Tech stacks:**
Golang, Gin, gRPC, MariaDB, GORM, Cobra, Pflag, Viper, Docker, GNU Make

## Try It Out!

1. Check prerequisites
  - [Docker 20.10+](https://docs.docker.com/get-docker/)
  - [GNU make](https://www.gnu.org/software/make/)

2. Check out the repository to your local machine and go to the project root directory.
    ```sh
    git clone git@github.com:josephzxy/timer_apiserver.git
    cd timer_apiserver
    ```

3. Build & launch with `docker compose`
    ```sh
    make docker.compose.up
    ```
4. Try out RESTful APIs & gRPC APIs
    ```sh
    # Create a timer named "hello"
    make demo.rest.post.hello

    # Display current content in the table
    make demo.db.show

    # Get the timer named "hello"
    make demo.rest.get.hello
    # Get all timers
    make demo.rest.get.all

    # Update the timer named "hello" to a new name "hello_again"
    make demo.rest.put.hello DEMO_REST_PUT_NAME=hello_again
    
    # Get all pending timers
    # Timers are pending if they are not deleted and not triggered yet
    make demo.grpc.getallpending

    # Delete the timer named "hello_again"
    make demo.rest.delete.hello_again
    ```

## Overview

[![System overview](docs/images/system_overview.svg)](https://drive.google.com/file/d/1B9L1sRXv4_FnJyslSze-C8GQ56jMQzyy/view?usp=sharing)

Timer API Server is a demo Golang project that implements RESTful APIs and gRPC APIs for Timer.

- Loose-coupling design where high-level pluggable controllers(e.g. REST API Server) share layered low-level resource managing service(e.g. MySQL middleware).
- Aligned with [12-factor-app](https://12factor.net/) methodology with supports for taking configurations from various sources, treating logs as event streams, launching the app as stateless processes, etc.
- Achieved [95% test coverage on CodeCov](https://app.codecov.io/gh/josephzxy/timer_apiserver) and [A+ on Go Report Card](https://goreportcard.com/report/github.com/josephzxy/timer_apiserver). Leverages GNU Make to automate trivial tasks. [Lints, tests, and builds on each commit with CircleCI](https://circleci.com/gh/josephzxy/timer_apiserver).
- Enabled automated building and serving with Docker Compose.

## API Doc
- **RESTful API**

    View swagger doc as a webpage by
    ```bash
    make swagger
    ```
    Or check out [swagger.yml](api/rest/swagger/swagger.yml) directly.

- **gRPC API**

    Please check out [timer.proto](api/grpc/timer.proto).
