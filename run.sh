#!/usr/bin/env bash
export GO_ENV="tests"

(! pidof go-api-jwt-v2) || sudo kill -9 $(pidof go-api-jwt-v2)

echo "GO_ENV="$GO_ENV
go build -o go-api-jwt-v2
./go-api-jwt-v2