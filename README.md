# ginDemo

## 一、项目内容

基于gin框架的demo项目,本项目基于B站视频Go语言 Gin+Vue 前后端分离实战 - OceanLearn
[https://www.bilibili.com/video/BV1CE411H7bQ?t=7](https://www.bilibili.com/video/BV1CE411H7bQ?t=7)

## 实现功能:
1.登录、注册、查看登录用户信息、文章分类的增删改查、文章发布的增删改查、文章分页 


2.中间件：（1）处理跨域问题 （2）用jwt加密用户登录信息检测用户登录状态 （3）捕获异常


3.数据库：mysql

**master**分支为后端go代码,**vue**分支为前端vue代码

## 二、怎样运行该项目

### 2.1 运行后端程序

> 先确保你电脑上正确安装了 golang 环境

从master分支拉取后端golang代码

```bash
# 拉取代码
git clone -b main git@github.com:EmperorEuler/ginDemo.git backend
# 进入项目目录
cd  backend
# 安装项目依赖
go get
```

打开 `config/application.yaml` 文件，修改数据库链接配置，修改项目运行端口，确保端口不被占用，参考如下

```yaml
server:
  port: 1016
datasource:
  driverName: mysql
  host: 127.0.0.1
  port: 3306
  database: goDemo
  username: root
  password: root
  charset: utf8
  loc: Asia/Guangzhou
```

启动项目

```bash
go run main.go
```

如果看到命令行终端输出以下路由信息，代表项目运行正常。如果不正常，检查一下数据库地址还有账号密码是否正确，同时确保运行的端口没有被占用

```bash
...
[GIN-debug] Listening and serving HTTP on :1016
```
