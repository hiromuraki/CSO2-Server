#!/bin/bash
set -e

podman run --rm -it \
    --name cso2-server \
    --userns keep-id \
    -p 1314:1314/tcp \
    -p 1315:1315/tcp \
    -p 30001:30001/tcp \
    -p 30002:30002/udp \
    ghcr.io/hm-gamesrv/cso2-server:latest