#!/usr/bin/env bash
export  GO_ENV="tests"
# export GOOS="linux" 
echo "GO_ENV="$GO_ENV
go build -o go-api-jwt
./go-api-jwt