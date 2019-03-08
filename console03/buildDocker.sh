#!/usr/bin/env bash

GOOS=linux GOARCH=amd64 go build -o console03_linux

#docker image rm -f console03_new

docker image build -t console03_new .


