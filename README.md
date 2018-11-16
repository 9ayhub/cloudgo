# cloudgo

这是一个简单的，类似 cloudgo 的应用。

## main.go
与老师的没多大差别。

## service.go
**使用了gin框架**。本来用的是martini框架，这是一个设计得很好的框架，但是存在一个比较严重的问题，就是反射使用太多导致效率过低（这个问题也导致了程序在访问量暴涨时内存上涨过快的问题），而且这个框架以经没有人维护了，而作者推荐使用风格很相近的gin框架，gin在github上的star数是最多的，而且仅仅从README看，其文档也是相当丰富的。

贴一下service.go的代码，十分的简单：
```
func NewServer(port string) { /*新建服务器*/
	r := gin.Default() /*获得路由实例*/
	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name") /*获取参数*/
		c.String(200, "Hello "+name) /*返回字符串*/
	})
	r.Run(":" + port) // listen and serve on 0.0.0.0:port
}
```

## 运行程序
```
go run main.go
```
![2](https://github.com/9ayhub/cloudgo/blob/master/pics/%5BGIN%5D.png)

访问localhost:8080

![1](https://github.com/9ayhub/cloudgo/blob/master/pics/1.png)

## curl测试

![3](https://github.com/9ayhub/cloudgo/blob/master/pics/curl.png)

## 压力测试

![4](https://github.com/9ayhub/cloudgo/blob/master/pics/ab1.png)
![5](https://github.com/9ayhub/cloudgo/blob/master/pics/ab2.png)
![6](https://github.com/9ayhub/cloudgo/blob/master/pics/ab3.png)
