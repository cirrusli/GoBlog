# GoBlog
Summer vacation

The first front-end and back-end separation project

## 环境依赖
go 1.18

MySQL 5.7

Redis x64-3.2.100

## 部署步骤
1. 启动Vue服务器
2. 启动Go服务器


## 目录结构描述
```
│  go.mod
│  go.sum
│  README.md
│          
├─cmd
│      main.go
│      
├─common
│      commonResponse.go
│      encoder.go
│      idGenerator.go
│      mysql.go
│      redis.go
│      sensitiveFilter.go
│      snowflakeid.go
│      
├─config
│      config.go
│      config.toml
│      
├─document
│      Question.md
│      sensitiveDict.txt
│      
└─internal
    ├─controller
    │      articleHandler.go
    │      commentHandler.go
    │      followHandler.go
    │      likeHandler.go
    │      userHandler.go
    │      
    ├─dao
    │      articleDAO.go
    │      commentDAO.go
    │      likeDAO.go
    │      userfollowDAO.go
    │      userinfoDAO.go
    │      
    ├─middleware
    │      authJWT.go
    │      middlewareSourceCode.go
    │      simpleImplementation.go
    │      
    ├─model
    │      article.go
    │      comment.go
    │      config.go
    │      like.go
    │      resmsg.go
    │      userfollow.go
    │      userinfo.go
    │      
    ├─router
    │      router.go
    │      
    └─service
            articleService.go
            commentService.go
            followService.go
            likeService.go
            userService.go
            
```
