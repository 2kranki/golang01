#!/usr/bin/env bash

GOOS=linux GOARCH=amd64 go build -o console04_linux

#docker image rm -f kranki/console04

docker image build -f Dockerfile.txt -t kranki/console04 .


