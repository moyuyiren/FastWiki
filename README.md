# FastWiki

#### 本系统用于快速搭建一个拥有二级菜单的Wiki系统，

    通过修改数据库内信息实现针对不同项目的wiki支持

#### 采用技术栈

本系统采用GOlang作为主要编程语言
采用mysql与elasticSearch作为数据库
redis作为缓存为系统提供主要数据的服务 
支持日志采集与归档
可通过RabbitMQ快捷通知短信系统为用户提供账号安全服务

- [x]  用户及权限管理
- [x]  菜单及缓存刷新
- [x]  Rabbit通知下游系统
- [x]  基础日志系统
- [ ]  Es主要查询

项目目录
.
├─.idea
│  └─inspectionProfiles
├─Config  配置文件
├─Controller  分发层
│  └─reciprocal 前端返回封装
├─Dao   数据层
│  ├─ElasticSearch
│  ├─MySql
│  └─Redis
├─Logger 日志管理
├─Logic  逻辑层
├─Middlewares  中间件
├─Model  
├─RabbitMQ  MQ
├─RouterS   路由及中间件管理
├─Setting   配置文件实例化及热部署实现
├─Token     登录管理
└─Utils     工具
