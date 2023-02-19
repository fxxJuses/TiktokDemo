# 部署

只需要部署点赞和评论的功能即可。

其中需要迁移的代码有：

```
--config
	|-- Config.go # 里面是Redis和RabbitMQ的配置
	|-- douyin.sql # 用于生成mysql数据字段的sql
	|-- nginx.conf # 这个应该是不需要的部署的，里面是vsftp相关的配置
	|-- redis.conf # 里面是redis的相关配置，需要将其中的 bind信息修改成部署端口信息

--controller # 控制层
	|-- CommentController.go # 里面是评论控制层代码
	|-- FavoriteController.go # 里面是点赞控制层代码

--dao # 数据层代码
	|-- InitDao.go # mysql初始化，
	|-- LikeDao.go # 点赞dao层
	|-- CommentDao.go # 评论dao层
	|-- VideoDao.go #视频dao层
	|-- UserDao.go # 用户到层

--mideelware #中间件代码
	|-- rabbitMQ
		|-- RabbitMQInit.go # rabbitMQ初始化
		|-- LikeMQ.go # 点赞队列
		|-- CommentMQ.go # 评论队列
	|-- redis
		|-- RedisInit.go # reids初始化操作
		|-- CommentDB.go # 点赞redis操作
		|-- UserDB.go # 用户Redis操作
		|-- VideoDb.go # 视频Redis操作

--models # entity代码
	|-- ....

--service # 服务层代码
	|--CommentServiceImpl.go # 评论实现层
	|--LikeServiceImpl.go # 点赞实现层
	|--VideoServiceImpl.go # 视频实现层
	|--UserServiceImpl.go # 用户实现层

-- utils # 工具代码
	|-- sensitiveFilter.go # 敏感词过滤器
	|-- SensitiveDict.txt # 敏感词库 
```

---------------

github链接：[fxxJuses/TiktokDemo (github.com)](https://github.com/fxxJuses/TiktokDemo)

团队项目链接：[seed30/TikTok at base (github.com)](https://github.com/seed30/TikTok/tree/base)

【注意实现】

一鸣提供的团队项目框架中，没有redis和RabbitMQ，需要将这两个环境配置，并添加到docker—compose里面，以便其他人开发。

开发平台为 Linux，我是在win上做的开发，可能会有一些细节需要注意。