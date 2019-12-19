#!/usr/bin/env bash

go env GO111MODULE=on
go env GOPROXY=https://goproxy.cn,direct

go build -o base-server main.go
chmod u+x ./restart.sh