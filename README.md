# GoBlog
Summer vacation front-end and back-end separation project

###########环境依赖
go 1.18
MySQL 5.7
Redis x64-3.2.100

###########部署步骤
1. 启动Vue服务器

2. 启动Go服务器


###########目录结构描述
├── README.md                   
├── cmd                         // 应用
│   └── main.go
├── common                  // 通用组件
│   ├── commonResponse.go
│   ├── encoder.go
│   ├── idGenerator.go
│   ├── mysql.go
│   ├── redis.go
│   ├── sensitiveFilter.go
│   └── snowflakeid.go
├── config                      
│   ├── config.go                // 配置控制
│   └── config.toml               // 配置
├── document                      
│   ├── Question.md               // 项目总结的问题
│   └── sensitiveDict               // 敏感词字典
├── internal
│   ├── controller               // 控制层
│   │   ├── articleHandler.go
│   │   ├── commentHandler.go
│   │   ├── userHandler.go
│   │   ├── likeHandler.go
│   │   └── followHandler.go
│   ├── dao                      // 数据库层
│   │   ├── articleDAO.go
│   │   ├── commentDAO.go
│   │   ├── userfollowDAO.go
│   │   └── userinfoDAO.go
│   ├── middleware               // 中间件
│   │   ├── authJWT.go
│   │   ├── middlewareSourceCode.go
│   │   └── simpleImplementation.go
│   ├── model                    // 模型层
│   │   ├── article.go
│   │   ├── comment.go
│   │   ├── like.go
│   │   ├── resmsg.go
│   │   ├── userinfo.go
│   │   └── userfollow.go
│   ├── router                      // 路由
│   │   └── router.go
│   ├── service                   // 服务层
│   │   ├── articleService.go
│   │   ├── commentService.go
│   │   └── userService.go
└── go.mod
    └──go.sum