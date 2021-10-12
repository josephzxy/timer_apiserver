#!/bin/bash
#
# The purpose of this script is to automate basic functional tests for RESTful & gRPC APIs of this project
# with human-vet needed outputs.
#
# Prerequisites:
# GNU Make - https://www.gnu.org/software/make/
# docker - https://docs.docker.com/get-docker/ 

set -euo pipefail

PROJECT_ROOT=$(cd $(dirname $(realpath ${0})); cd ../..; pwd -P)
OS=$(go env GOOS)
ARCH=$(go env GOARCH)

function run_make() {
    cd ${PROJECT_ROOT}; make $1
}

function run_test() {
    run_make demo.rest.healthz
    run_make demo.rest.post.hello
    run_make demo.db.show

    run_make demo.rest.get.hello
    run_make demo.rest.get.all
    run_make "demo.rest.put.hello DEMO_REST_PUT_NAME=hello_again"
    run_make demo.db.show

    run_make demo.grpc.getallpending
    run_make demo.rest.delete.hello_again
    run_make demo.db.show
}

function cleanup() {
    set +e
    run_make docker.compose.down > /dev/null 2>&1
    run_make mysql.docker.stop > /dev/null 2>&1

    for port in 8081 8082; do
        kill $(lsof -ti tcp:${port}) > /dev/null 2>&1
    done
    set -e
    run_make clean > /dev/null 2>&1
}

function with_env::local() {
    echo "=======> Test running in local environment"

    cleanup
    run_make build
    run_make mysql.docker.start
    BIN="${PROJECT_ROOT}/_output/build/${OS}/${ARCH}/apiserver"
    $BIN --config config/example.yml > /dev/null 2>&1 &
    sleep 1

    $1
    cleanup
}

function with_env::docker_compose() {
    echo "=======> Test running with docker compose"
    cleanup
    run_make docker.compose.up
    $1
    cleanup
}

function main() {
    with_env::local run_test
    with_env::docker_compose run_test
}

main
