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

### 压力测试参数解释：
...
Time taken for tests:   0.466 seconds   ###总请求时间
Complete requests:      1000     ###总请求数
Failed requests:        0     ###失败的请求数
**Requests per second:**    2146.38 [#/sec] (mean)      ###平均每秒的请求数
**Time per request:**      46.590 [ms] (mean)     ###平均每个请求消耗的时间
**Time per request:**       0.466 [ms] (mean, across all concurrent requests)  ###上面的请求除以并发数
Transfer rate:          268.30 [Kbytes/sec] received   ###传输速率
...
Percentage of the requests served within a certain time (ms)
  50%   37   ###50%的请求都在37Ms内完成
...

重点关注吞吐率（Requests per second）、用户平均请求等待时间（Time per request）指标：

**1、吞吐率（Requests per second）：**

服务器并发处理能力的量化描述，单位是reqs/s，指的是在某个并发用户数下单位时间内处理的请求数。某个并发用户数下单位时间内能处理的最大请求数，称之为最大吞吐率。

记住：吞吐率是基于并发用户数的。这句话代表了两个含义：

a、吞吐率和并发用户数相关

b、不同的并发用户数下，吞吐率一般是不同的

计算公式：总请求数/处理完成这些请求数所花费的时间，即

Request per second=Complete requests/Time taken for tests

必须要说明的是，这个数值表示当前机器的整体性能，值越大越好。


**2、用户平均请求等待时间（Time per request）：**

计算公式：处理完成所有请求数所花费的时间/（总请求数/并发用户数），即：

Time per request=Time taken for tests/（Complete requests/Concurrency Level）


**3、服务器平均请求等待时间（Time per request:across all concurrent requests）：**

计算公式：处理完成所有请求数所花费的时间/总请求数，即：

Time taken for/testsComplete requests

可以看到，它是吞吐率的倒数。

同时，它也等于用户平均请求等待时间/并发用户数，即

Time per request/Concurrency Level。


