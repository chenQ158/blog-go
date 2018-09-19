package main

import (
	"blog-new/models"
	_ "blog-new/routers"
	"blog-new/service"
	"blog-new/utils"
	log "code.google.com/log4go"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"
)

var etcds = []string{"http://192.168.91.130:2379"}

//func initConfig() {
//	service.InitSelector(etcds)
//	hystrixStreamHandler := hystrix.NewStreamHandler()
//	hystrixStreamHandler.Start()
//	// 监听
//	go http.ListenAndServe(net.JoinHostPort("172.16.1.60", "8090"), hystrixStreamHandler)
//}

func init() {
	// 初始化redis
	utils.InitClusterClient()

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	utils.InitConfig(etcds)
	resp, err := utils.GetKeyList("conf/authCenterRpcConf")
	if err != nil {
		log.Error(err)
	} else {
		log.Debug(resp)
	}

	models.RegisterDBUsingGorm()

	var FilterUser = func(ctx *context.Context) {
		// 如果是登录或者静态文件直接跳过
		url := ctx.Request.RequestURI
		// fmt.Println("url:",url)
		if exist := strings.Contains(url, "/login") || strings.Contains(url, "/static") || strings.Contains(url, "/swagger"); exist {
			return
		}

		token,ok := ctx.GetSecureCookie("chen", "Token")
		// fmt.Println("token,ok: ",token,ok)
		if !ok {
			ctx.Redirect(302,"/login")
			return
		}

		userptr := service.GetSessionUser(token)
		if userptr == nil {
			ctx.Redirect(302, "/login")
			return
		}
	}

	beego.InsertFilter("/user/*", beego.BeforeRouter, FilterUser)

}


func main() {
	//initConfig()
	beego.Run()
}


