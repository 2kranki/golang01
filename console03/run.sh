#!/usr/bin/env bash

docker container rm -f console03_1

docker container run --name console03_1 -p9000:9000 -d console03_new


