# GoZero-IM: 基于 Go-Zero 的分布式即时通讯系统

## 📌 项目简介
GoZero-IM 是一个基于 [Go-Zero](https://github.com/zeromicro/go-zero) 框架的高性能分布式即时通讯（IM）系统，支持**私聊、群聊、消息存储、离线推送**等核心功能，适用于高并发场景。

## 🎯 主要特性
- **高性能架构**：采用 Go-Zero 微服务架构，支持大规模并发连接。
- **多端支持**：支持 **Web、iOS、Android、小程序** 等多端接入。
- **消息可靠性**：使用 Redis 进行消息存储与离线消息推送，确保消息不丢失。
- **分布式部署**：基于 Kubernetes + Docker 进行容器化部署，支持水平扩展。
- **完整 API**：提供 Restful API，方便对接其他系统。

## 🏗️ 技术栈
- **后端**：Golang + Go-Zero
- **数据库**：MySQL
- **缓存**：Redis
- **消息队列**：Kafka / RabbitMQ（可选）
- **协议支持**：WebSocket / gRPC / HTTP
- **部署**：Docker + Kubernetes + ArgoCD

## 📂 项目结构
```plaintext
GoZero-IM/
├── api          # API 定义
├── app          # 业务逻辑层
├── cmd          # 启动入口
├── config       # 配置文件
├── docs         # 项目文档
├── internal     # 内部实现
│   ├── logic    # 业务逻辑
│   ├── model    # 数据模型
│   ├── svc      # 服务上下文
│   ├── handler  # 路由处理
├── pkg          # 工具包
├── scripts      # 启动/部署脚本
└── test         # 测试代码
```

## 🚀 快速开始
### 1️⃣ 环境准备
请确保本地已安装：
- **Go 1.18+**
- **MySQL 8.0+**
- **Redis 6.0+**
- **Docker & Kubernetes（可选）**

### 2️⃣ 数据库初始化
```sql
CREATE DATABASE gozero_im;
USE gozero_im;

-- 创建用户表
CREATE TABLE users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建消息表
CREATE TABLE messages (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    sender_id BIGINT NOT NULL,
    receiver_id BIGINT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 3️⃣ 启动服务
```bash
# 克隆项目
git clone https://github.com/your-repo/gozero-im.git
cd gozero-im

# 安装依赖
go mod tidy

# 运行服务
go run cmd/main.go -f config.yaml
```

## 📡 API 使用
### 1️⃣ 用户注册
```http
POST /api/register
Content-Type: application/json

{
    "username": "user1",
    "password": "password123"
}
```

### 2️⃣ 发送消息
```http
POST /api/message/send
Content-Type: application/json

{
    "sender_id": 1,
    "receiver_id": 2,
    "content": "Hello, how are you?"
}
```

### 3️⃣ 获取历史消息
```http
GET /api/message/history?user_id=1&friend_id=2
```

## ⚙️ Docker 部署
```bash
# 构建 Docker 镜像
docker build -t gozero-im .

# 运行容器
docker run -d -p 8080:8080 --name gozero-im gozero-im
```

## ☁️ Kubernetes 部署
```bash
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

## 🎯 贡献指南
欢迎任何形式的贡献！
1. **Fork** 本项目
2. 创建你的新分支 (`git checkout -b feature-xxx`)
3. 提交你的修改 (`git commit -m 'add feature xxx'`)
4. Push 到你的分支 (`git push origin feature-xxx`)
5. 提交 Pull Request

## 📄 许可证
本项目采用 **MIT 许可证**，详情请查阅 LICENSE 文件。

---
📧 **联系我们**：如果有任何问题，欢迎提交 Issue 或加入我们的开源交流群！

