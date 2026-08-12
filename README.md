# CSO2 Server

Counter-Strike Online 2 服务端。用于 **2017 年国服** 和 **2018 年韩服** 客户端。

> *Counter-Strike Online 2 归 NEXON 所有，本项目仅用于学习之用。*

## 1. 关于本项目

本项目是 [LoveBeforT/CSO2-Server](https://github.com/LoveBeforT/CSO2-Server) 的改进版本，主要改进：

- **纯 Go 静态编译**：移除 cgo/iconv 依赖，编码转换使用 `golang.org/x/text` 纯 Go 实现
- **优化的容器化运行**：支持 Docker / Podman 一行命令启动
- **GitHub Actions 自动构建**：推送镜像到 GHCR，无需本地编译

## 2. 快速开始（Docker / Podman）

### 2.1. 从镜像仓库拉取

```bash
# Docker
docker pull ghcr.io/<owner>/cso2-server:latest

# Podman
podman pull ghcr.io/<owner>/cso2-server:latest
```

### 2.2. 运行

```bash
# Docker
docker run -d --name cso2-server \
    -p 1314:1314 \
    -p 1315:1315 \
    -p 30001:30001 \
    -p 30002:30002/udp \
    -v "$(pwd)/database:/app/CSO2-Server/database" \
    ghcr.io/<owner>/cso2-server:latest

# Podman
podman run -d --name cso2-server \
    -p 1314:1314 \
    -p 1315:1315 \
    -p 30001:30001 \
    -p 30002:30002/udp \
    -v "$(pwd)/database:/app/CSO2-Server/database" \
    ghcr.io/<owner>/cso2-server:latest
```

### 2.3. 本地构建

```bash
git clone https://github.com/<owner>/CSO2-Server.git
cd CSO2-Server
docker build -t cso2-server .
```

## 3. 端口说明

| 端口  | 协议 | 用途                                                |
| ----- | ---- | --------------------------------------------------- |
| 1314  | TCP  | Web 注册页面（浏览器访问 `http://<服务器IP>:1314`） |
| 1315  | TCP  | GM 控制台                                           |
| 30001 | TCP  | 游戏主连接                                          |
| 30002 | UDP  | 联机通信                                            |

## 4. 注册与登录

1. 启动服务端后，浏览器访问 `http://<服务器IP>:1314`，点击右上角 **Register** 注册账号（注：默认无需邮箱和验证码）
2. 客户端启动器连接时指定服务端 IP 即可登录

## 5. 数据持久化

玩家数据保存于容器中的目录 `/app/CSO2-Server/database`

## 7. 原项目

README 与更多细节见原始项目：[LoveBeforT/CSO2-Server](https://github.com/LoveBeforT/CSO2-Server)
