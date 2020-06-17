这是我初学golang之后第一个用的框架，跟着[煎鱼大佬学的](https://github.com/EDDYCJY/go-gin-example),按部就班，收获了非常多，但之前也在Python中用过`Django`，`Flask`等框架，学习的过程中也修改了一些东西：


1. 修改数据库数据，默认为[mysql的官方示例数据](https://www.yiibai.com/mysql/sample-database.html)，数据多，好操作
2. 把Post请求数据的格式换为`application/json`，这个更好理解好操作
3. JWT认证我使用了[第三方库](github.com/appleboy/gin-jwt/v2)作为中间件
4. Response信息我遵循了rest风格，并不是所有请求都返回状态200，错误信息也简单做了修改，答题不变，具体看[这边](https://tech.crandom.com/post/2019/restful-status-code/)
5. 增加了`docker-compose`配置以及相应的`Makefile`配置
6. 增加了`api_test`文件作为单元测试
7. 修改了部分配置文件名，遵循golint，如`HttpPort`改为`HTTPPort`
8. 暂时没加log以及redis缓存，有空研究


### TODO
1. Redis 缓存
2. gin日志系统


### 注意

1. `docker-compose`配置文件夹为compose文件夹，额外的配置文件，数据库源文件等
2. ！！！ `docker-compose`第一次启动 web以及nginx都会不成功，因为db服务会导入数据库源文件，因为里面数据记录较多需要等个半分钟，再手动启动web以及nginx服务，这只会在第一次的时候发生
3. docker这边测试端口为8000，本地我改为了5000
4. swagger为http://127.0.0.1:8000/swagger/index.html
5. 如果不用docker必须自带mysq并导入相关数据


运行的话直接 `docker-compose up -d`

## 6.17更新

使用了官方示例的log配置
```python
// Disable Console Color, you don't need console color when writing the logs to file.
gin.DisableConsoleColor()

// Logging to a file.
f, _ := os.Create("gin.log")
gin.DefaultWriter = io.MultiWriter(f)
```
但只有请求信息，自己的`log.println`没法存储到文件里