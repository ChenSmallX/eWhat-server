# eWhat-server
the back-end of the WX-program:="今天吃什么". Coding by Golang...

微信小程序“今天吃什么”后台程序，使用Go语言编写,包括数据请求接口（简单路由），数据库控制器，日志控制器三部分。

## 项目分层结构
1. WebRouter
	- 负责开启端口监听
	- 负责解析请求（GET&POST）
	- 负责调用SQLController对数据库进行操作
	- 与Logger交互生成[请求、获取等操作的]日志（不同日志等级）
2. SQLController
	- 负责操作数据库（MySQL）
	- 与Logger交互生成[数据库调用的]日志
3. Logger
	- 被两个模块调用，生成日志文件         日志模块

## 结构图简画
internet -> { [ WebRouter -> SQLController -> (SQL)(SeparatedSoftware) ] -> Logger }(WholeSystem)

## TODO
- [ ] WebRouter
    - [x] 和前端工程师对接口
	- [ ] GET Class method
	- [ ] POST Class method
- [ ] SQLController
	- [ ] SQL-Builder
	- [ ] SQL-manipulator(增删改查)
- [ ] Logger
	- [ ] file-manipulator
	- [ ] formatter
	- [ ] filter
	- [ ] generator

## WebRouter

本模块基于Golang编写的ECHO包编写，ECHO具有编写简单，没有太多复杂功能的特点，适合用于本项目“数据传输仅通过JSON和图片”的应用场景。

## SQLController

本模块基于Golang编写的GO-MySQL-Driver包编写，数据库程序使用MySQL

## Logger

本模块基于Golang官方log包编写

