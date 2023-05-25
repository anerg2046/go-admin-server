# GO-ADMIN-SERVER 

[![license](https://img.shields.io/github/license/anerg2046/go-admin-server.svg)](LICENSE)

## 简介

本项目是基于 [go-app-template](https://github.com/anerg2046/go-app-template) 开发的通用后台基础代码

配套的前端代码为 [go-admin-front](https://github.com/anerg2046/go-admin-front)

### 已完成的部分：

- 完善的菜单管理
- 完善的角色管理
- 完善的用户管理

### 特性

- 基于casbin对菜单的显示，和后端请求接口均可按角色分配权限

### 说明

采用了git子模块的方式，所以clone的时候请用下面的命令

`git clone --recurse-submodules git@github.com:anerg2046/go-admin-server.git`

需要配置的东西全部在 `config` 目录中，还请自行查看

### 演示地址

https://admin.fabraze.com/

超级管理员：

admin   admin1234

普通用户：

test    test123456

### 其它

克隆代码以后，运行`go mod tidy`安装依赖

在根目录有两个快捷命令

- `make db` 创建数据库，和默认数据，前提是已经配置好数据库
- `make api` 执行服务端接口程序