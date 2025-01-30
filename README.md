# Pet 项目

一个基于 Go 语言开发的宠物管理系统后端服务。

## 项目架构

本项目采用清晰的分层架构:

- `cmd`: 应用程序入口
- `configs`: 配置文件
- `internal`: 内部代码
  - `route`: HTTP 路由及处理器
  - `service`: 业务逻辑层
  - `cron`: 定时任务
- `log`: 日志文件

## 技术栈

- Web 框架: [Gin](https://github.com/gin-gonic/gin)
- 依赖注入: [Wire](https://github.com/google/wire)
- 日志: [Zap](https://github.com/uber-go/zap)

## 主要功能

- 用户管理系统
- RESTful API
- 优雅关机
- 跨域支持
- 日志记录

## 运行

### 环境要求

- Go 1.16+
- MySQL 5.7+
