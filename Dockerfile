FROM golang:1.26-alpine AS builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0

WORKDIR /build

# 先拷贝依赖清单，利用 Docker 层缓存
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# -tags timetzdata：嵌入时区数据库（alpine 无 /usr/share/zoneinfo）
# -ldflags '-s -w'：去掉符号表，减小体积
RUN go build -tags timetzdata -ldflags '-s -w' -o /out/CSO2-Server-bin .

# exe 在 /app，资源必须在 /app/CSO2-Server/{assert,locales,database,configure}
RUN mkdir -p /out/CSO2-Server/database \
             /out/CSO2-Server/configure \
    && cp -r ./assert ./locales /out/CSO2-Server/ \
    && cp ./configure/server.conf /out/CSO2-Server/configure/

FROM alpine:latest

EXPOSE 1314/tcp 1315/tcp 30001/tcp 30002/udp

RUN addgroup -g 1000 gamesrv \
    && adduser -u 1000 -D -G gamesrv -s /bin/sh gamesrv \
    && mkdir -p /app \
    && chown -R 1000:1000 /app

COPY --from=builder --chown=1000:1000 /out /app

# 玩家数据持久化（json 存档目录），运行时可挂载到宿主机
VOLUME ["/app/CSO2-Server/database"]

USER 1000:1000
WORKDIR /app
CMD ["/app/CSO2-Server-bin"]
