# ⏳ Timer API Server
[![codecov](https://codecov.io/gh/josephzxy/timer_apiserver/branch/develop/graph/badge.svg?token=AU3193CPC3)](https://codecov.io/gh/josephzxy/timer_apiserver)

⏳ Timer API server is a demo Golang project that implements both external-facing RESTful APIs and internal-facing gRPC APIs for managing RESTful resource Timer.

**Major Tech stacks:**
Golang, Gin, gRPC, MariaDB, GORM, Cobra, pflag, viper, Docker, GNU make

## Try It Out!

1. Check prerequisites
  - [Docker 20.10+](https://docs.docker.com/engine/release-notes/#version-2010)
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
4. Try out RESTful APIs
    ```sh
    # Display current content in the table
    make demo.db.show

    # Create a timer named "hello"
    make demo.rest.post.hello

    # Get the timer named "hello"
    make demo.rest.get.hello
    # Get all timers
    make demo.rest.get.all

    # Update the timer named "hello" to a new name "hello_again"
    make demo.rest.put.hello DEMO_REST_PUT_NAME=hello_again
    
    # Delete the timer named "hello"
    make demo.rest.delete.hello
    ```
5. Try out gRPC APIs
    ```sh
    # Get all pending timers
    # Timers are pending if they are not deleted and not triggered yet
    make demo.grpc.getallpending
    ```

## RESTful API Doc

View swagger doc as a webpage by
```bash
make swagger
```
Or check out [swagger.yml](api/rest/swagger/swagger.yml) directly.

## gRPC API Doc
Please check out [timer.proto](api/grpc/timer.proto).
