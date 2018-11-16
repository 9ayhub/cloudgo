package service

import (
	"github.com/go-martini/martini" /*使用martini*/
)

func NewServer(port string) { /*新建服务器*/
	m := martini.Classic()
	/*添加参数[name]martini的参数中*/
	m.Get("/hello/:name", func(params martini.Params) string {
		return "Hello " + params["name"] + " :)"
	})
	/*对应main函数中的端口*/
	m.RunOnAddr(":" + port)
}
