# CSO2 Server

Counter-Strike Online 2 服务端。用于 **2017 年国服** 和 **2018 年韩服** 客户端。

> *Counter-Strike Online 2 归 NEXON 所有，本项目仅用于学习之用。*

## 1. 简述

本项目是 [LoveBeforT/CSO2-Server](https://github.com/LoveBeforT/CSO2-Server) 的改进版本，主要改进：

- **纯 Go 静态编译**：移除 cgo/iconv 依赖，编码转换使用 `golang.org/x/text` 纯 Go 实现
- **优化的容器化运行**：支持 Docker / Podman 一行命令启动
- **GitHub Actions 自动构建**：推送镜像到 GHCR，无需本地编译

**可用版本：**

| 游戏模式   | 镜像 tag |
| ---------- | -------- |
| 最新构建版 | `latest` |

## 2. 资源占用信息

### 2.1. 端口

| 端口号 | 协议 | 说明         |
| ------ | ---- | ------------ |
| 1314   | TCP  | Web 注册页面 |
| 1315   | TCP  | GM 控制台    |
| 30001  | TCP  | 游戏主连接   |
| 30002  | UDP  | 联机通信     |

### 2.2. 持久卷

| 容器路径                    | 说明     |
| --------------------------- | -------- |
| `/app/CSO2-Server/database` | 玩家数据 |

## 3. 快速开始（Docker / Podman）

### 3.1. 构建并运行（Docker）

```bash
docker build -t cso2:temp . && \
    docker run --rm -it \
        -p 1314:1314 \
        -p 1315:1315 \
        -p 30001:30001 \
        -p 30002:30002/udp \
        -v ./database:/app/CSO2-Server/database \
        cso2:temp
```

### 3.2. 运行服务器（Podman）

```bash
IMAGE=ghcr.io/hm-gamesrv/cso2:latest

if ! podman pull "$IMAGE"; then
    exit 1
fi

podman run --rm -it \
    --name cso2 \
    --userns keep-id \
    --network pasta \
    -p 1314:1314 \
    -p 1315:1315 \
    -p 30001:30001 \
    -p 30002:30002/udp \
    -v ./database:/app/CSO2-Server/database \
    "$IMAGE"
```

## 4. 注册与登录

1. 启动服务端后，浏览器访问 `http://<服务器IP>:1314`，点击右上角 **Register** 注册账号（注：默认无需邮箱和验证码）
2. 客户端启动器连接时指定服务端 IP 即可登录

## 5. 原项目

README 与更多细节见原始项目：[LoveBeforT/CSO2-Server](https://github.com/LoveBeforT/CSO2-Server)
