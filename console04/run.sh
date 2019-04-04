#!/usr/bin/env bash

docker container rm -f console04_1

docker container run --name console04_1 -p8080:8080 -d kranki/console04


