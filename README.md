```markdown
基于Go语言开发新闻爬虫系统:
    --简单爬虫
    --并发爬虫
    --分布式爬虫
```

```markdown
分布式系统消息传递的方法：REST、RPC、中间件
    对外：REST
    模块内部：RPC
    模块之间：中间件、REST
```

```markdown
分布式架构 VS 微服务架构
    分布式架构：指导节点之间如何通信
    微服务架构：鼓励按业务划分模块
    微服务架构通过分布式架构来实现

多层架构 VS 微服务架构
    微服务架构具有更多的“服务”
```

```markdown
1、限流问题
    单节点能够承受的流量有限，将 worker 放到不同的节点 （不同的服务器）
2、去重问题
    单节点能够承受的去重数据量有限
    无法保存之前（重启）的去重结果
    基于Key-Value Store(如Redis) 进行分布式去重
3、数据存储问题
    存储模块独立为一个服务
```