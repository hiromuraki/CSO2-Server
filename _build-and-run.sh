#!/bin/bash
set -e

docker build -t cso2-server:temp . &&\
    docker run --rm -it \
        -p 1314:1314/tcp \
        -p 1315:1315/tcp \
        -p 30001:30001/tcp \
        -p 30002:30002/udp \
        cso2-server:temp