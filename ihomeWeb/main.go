package main

import (
	"net/http"
	"sss/ihomeWeb/handler"

	"github.com/julienschmidt/httprouter"

	"github.com/micro/go-micro/util/log"

	"github.com/micro/go-micro/web"
	// 只运行一次，用来构建数据库结构
	_ "sss/ihomeWeb/models"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.IhomeWeb"),
		web.Version("latest"),
		web.Address(":8999"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}
	// 使用restful路由插件替代原有的路由
	rou := httprouter.New()
	// 映射静态页面
	rou.NotFound = http.FileServer(http.Dir("html"))
	// 获取地区请求
	rou.GET("/api/v1.0/areas", handler.GetArea)
	// 获取session请求
	rou.GET("/api/v1.0/session", handler.GetSession)
	// 提交登录请求
	rou.POST("/api/v1.0/session", handler.PostSession)
	// 退出登录请求
	rou.DELETE("/api/v1.0/session", handler.DeleteSession)
	// 获取首页轮播请求
	rou.GET("/api/v1.0/house/index", handler.GetIndex)
	// 获取验证码图片
	rou.GET("/api/v1.0/imagecode/:uuid", handler.GetImageCode)
	// 获取短信验证码
	rou.GET("/api/v1.0/smscode/:mobile", handler.GetSmsCode)
	// 提交注册表单
	rou.POST("/api/v1.0/users", handler.PostReg)
	// 获取用户信息
	rou.GET("/api/v1.0/user", handler.GetUserInfo)
	// 上传用户头像
	rou.POST("/api/v1.0/user/avatar", handler.PostAvatar)

	// 注册服务
	service.Handle("/", rou)

	// service.Handle("/", http.FileServer(http.Dir("html")))
	// service.HandleFunc("/IhomeWeb/call", handler.IhomeWebCall)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
